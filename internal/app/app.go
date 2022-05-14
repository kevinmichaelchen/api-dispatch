package app

import (
	"database/sql"
	"github.com/kevinmichaelchen/api-dispatch/internal/distance"
	"github.com/kevinmichaelchen/api-dispatch/internal/service"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
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

type ServiceParams struct {
	fx.In
	GRPCServer      *grpc.Server
	DB              *sql.DB
	DistanceService *distance.Service `optional:"true"`
}

func NewService(p ServiceParams) *service.Service {
	return service.NewService(p.DB, p.DistanceService)
}
