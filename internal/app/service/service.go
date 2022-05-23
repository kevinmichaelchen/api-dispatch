package service

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kevinmichaelchen/api-dispatch/internal/service"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/db"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/geo"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/fx"
	"go.uber.org/zap"
	gmaps "googlemaps.github.io/maps"
	"os"
)

var Module = fx.Module("service",
	fx.Provide(
		NewService,
		NewGeoService,
		NewMapsClient,
		NewDataStore,
	),
)

type ServiceParams struct {
	fx.In
	DataStore       *db.Store
	DistanceService *geo.Service
}

func NewService(p ServiceParams) *service.Service {
	return service.NewService(p.DataStore, p.DistanceService)
}

func NewDataStore(sqlDB *sql.DB, redisClient *redis.Client) *db.Store {
	return db.NewStore(sqlDB, redisClient)
}

func NewMapsClient(logger *zap.Logger) (*gmaps.Client, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		logger.Warn("Missing API Key for Google Maps... Initializing in degraded state...")
		return nil, nil
	}
	c, err := gmaps.NewClient(
		gmaps.WithAPIKey(apiKey),
		gmaps.WithHTTPClient(otelhttp.DefaultClient),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build Google Maps client: %w", err)
	}
	return c, nil
}

type GeoServiceParams struct {
	fx.In
	GoogleClient *gmaps.Client `optional:"true"`
}

func NewGeoService(logger *zap.Logger, p GeoServiceParams) (*geo.Service, error) {
	return geo.NewService(p.GoogleClient, otelhttp.DefaultClient), nil
}
