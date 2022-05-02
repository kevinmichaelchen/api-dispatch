package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/kevinmichaelchen/api-dispatch/internal/distance"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/service"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"googlemaps.github.io/maps"
	"log"
	"net"
	"os"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")
	return logger
}

func NewGRPCServer(lc fx.Lifecycle, logger *log.Logger) (*grpc.Server, error) {
	// TODO configure options here
	//var opts grpc.ServerOption
	s := grpc.NewServer()
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			logger.Print("Starting gRPC server.")
			// TODO make configurable
			address := fmt.Sprintf(":%d", 8080)
			lis, err := net.Listen("tcp", address)
			if err != nil {
				logger.Fatalf("Failed to listen on address \"%s\": %v", address, err)
			}
			go s.Serve(lis)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping gRPC server.")
			// GracefulStop stops the gRPC server gracefully. It stops the server from
			// accepting new connections and RPCs and blocks until all the pending RPCs are
			// finished.
			s.GracefulStop()
			return nil
		},
	})
	return s, nil
}

func NewDatabase(logger *log.Logger, lc fx.Lifecycle) (*sql.DB, error) {
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
			logger.Print("Closing DB connection.")
			err := db.Close()
			if err != nil {
				logger.Printf("Failed to close DB connection: %v\n", err)
				return err
			}
			logger.Println("Successfully closed DB connection")
			return err
		},
	})
	return db, err
}

type ServiceParams struct {
	fx.In
	Logger          *log.Logger
	GRPCServer      *grpc.Server
	DB              *sql.DB
	DistanceService *distance.Service `optional:"true"`
}

func NewService(p ServiceParams) *service.Service {
	return service.NewService(p.Logger, p.DB, p.DistanceService)
}

func NewMapsClient() (*maps.Client, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, errors.New("missing API_KEY for Google Maps")
	}
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewDistanceService(logger *log.Logger, client *maps.Client) (*distance.Service, error) {
	if client == nil {
		return nil, errors.New("no maps client")
	}
	return distance.NewService(client), nil
}

// Register mounts our HTTP handler on the mux.
//
// Register is a typical top-level application function: it takes a generic
// type like *grpc.Server, which typically comes from a third-party library,
// and introduces it to a type that contains our application logic. In this
// case, that introduction consists of registering a gRPC handler. Other typical
// examples include registering HTTP servers and starting queue consumers.
//
// Fx calls these functions invocations, and they're treated differently from
// the constructor functions above. Their arguments are still supplied via
// dependency injection, and they may still return an error to indicate
// failure, but any other return values are ignored.
//
// Unlike constructors, invocations are called eagerly. See the main function
// below for details.
func Register(server *grpc.Server, svc *service.Service) {
	v1beta1.RegisterDispatchServiceServer(server, svc)
	reflection.Register(server)
}
