package kafka

// SIGUSR1 toggle the pause/resume consumption
import (
	"context"
	"github.com/Shopify/sarama"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// Sarama configuration options
var (
	brokers         = "127.0.0.1:9092"
	version         = "3.1.1"
	consumerGroupID = "api-dispatch"
	topics          = "driver-locations"
	assignor        = "range" // Consumer group partition assignment strategy (range, roundrobin, sticky)
	oldest          = true
	verbose         = false
)

var Module = fx.Module("kafka",
	fx.Provide(
		NewConsumerGroup,
	),
)

// NewConsumerGroup creates a new Kafka consumer group.
// Inspired by:
// https://github.com/Shopify/sarama/blob/main/examples/consumergroup/main.go
// https://github.com/confluentinc/cp-all-in-one/blob/7.1.0-post/cp-all-in-one-community/docker-compose.yml#L4-L34
func NewConsumerGroup(logger *zap.Logger, lc fx.Lifecycle) (sarama.ConsumerGroup, error) {
	keepRunning := true
	logger.Info("Starting a new Sarama consumer")
	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		return nil, err
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version

	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	/**
	 * Setup a new Sarama consumer group
	 */
	consumer := Consumer{
		ready:  make(chan bool),
		logger: logger,
	}

	var client sarama.ConsumerGroup

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			cc, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), consumerGroupID, config)
			if err != nil {
				return err
			}
			client = cc

			consumptionIsPaused := false
			wg := &sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					// `Consume` should be called inside an infinite loop, when a
					// server-side rebalance happens, the consumer session will need to be
					// recreated to get the new claims
					if err := client.Consume(ctx, strings.Split(topics, ","), &consumer); err != nil {
						logger.Panic("Error from consumer", zap.Error(err))
					}
					// check if context was cancelled, signaling that the consumer should stop
					if ctx.Err() != nil {
						return
					}
					consumer.ready = make(chan bool)
				}
			}()

			<-consumer.ready // Await till the consumer has been set up
			logger.Info("Sarama consumer up and running!...")

			sigusr1 := make(chan os.Signal, 1)
			signal.Notify(sigusr1, syscall.SIGUSR1)

			for keepRunning {
				select {
				case <-ctx.Done():
					logger.Info("terminating: context cancelled")
					keepRunning = false
				case <-sigusr1:
					toggleConsumptionFlow(logger, client, &consumptionIsPaused)
				}
			}
			wg.Wait()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Closing Kafka connection...")
			err := client.Close()
			if err != nil {
				logger.Error("Failed to close Kafka connection", zap.Error(err))
			} else {
				logger.Info("Successfully closed Kafka connection")
			}
			return err
		},
	})

	return client, nil
}

func toggleConsumptionFlow(logger *zap.Logger, client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		logger.Info("Resuming consumption")
	} else {
		client.PauseAll()
		logger.Info("Pausing consumption")
	}

	*isPaused = !*isPaused
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready  chan bool
	logger *zap.Logger
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		// TODO convert bytes to protobuf
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}

	return nil
}
