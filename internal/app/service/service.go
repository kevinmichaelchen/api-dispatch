package service

import (
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/kevinmichaelchen/api-dispatch/internal/service"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/db"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/distance"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"googlemaps.github.io/maps"
	"os"
)

var Module = fx.Module("service",
	fx.Provide(
		NewService,
		NewDistanceService,
		NewMapsClient,
		NewDataStore,
	),
)

type Params struct {
	fx.In
	DataStore       *db.Store
	DistanceService *distance.Service `optional:"true"`
}

func NewService(p Params) *service.Service {
	return service.NewService(p.DataStore, p.DistanceService)
}

func NewDataStore(sqlDB *sql.DB) *db.Store {
	return db.NewStore(sqlDB)
}

func NewMapsClient() (*maps.Client, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, errors.New("missing API_KEY for Google Maps")
	}
	c, err := maps.NewClient(
		maps.WithAPIKey(apiKey),
		maps.WithHTTPClient(otelhttp.DefaultClient),
	)
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