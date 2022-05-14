package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/kevinmichaelchen/api-dispatch/internal/distance"
	"github.com/kevinmichaelchen/api-dispatch/internal/service"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"googlemaps.github.io/maps"
	"log"
	"os"
)

func NewLogger() *zap.Logger {
	// TODO configure log options
	logger, err := zap.NewProduction(
		zap.AddCaller(),
	)
	if err != nil {
		log.Fatalf("failed to build zap logger: %v", err)
	}
	return logger
}

func NewDatabase(logger *zap.Logger, lc fx.Lifecycle) (*sql.DB, error) {
	// TODO make configurable
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/dispatch?sslmode=disable")
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Closing DB connection.")
			err := db.Close()
			if err != nil {
				logger.Error("Failed to close DB connection", zap.Error(err))
				return err
			}
			logger.Info("Successfully closed DB connection")
			return err
		},
	})
	return db, err
}

type ServiceParams struct {
	fx.In
	GRPCServer      *grpc.Server
	DB              *sql.DB
	DistanceService *distance.Service `optional:"true"`
}

func NewService(p ServiceParams) *service.Service {
	return service.NewService(p.DB, p.DistanceService)
}

func NewMapsClient() (*maps.Client, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, errors.New("missing API_KEY for Google Maps")
	}
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to build Google Maps client: %w", err)
	}
	return c, nil
}

func NewDistanceService(logger *zap.Logger, client *maps.Client) (*distance.Service, error) {
	if client == nil {
		return nil, errors.New("no maps client")
	}
	return distance.NewService(client), nil
}
