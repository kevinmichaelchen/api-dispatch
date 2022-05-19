package service

import (
	"context"
	"github.com/kevinmichaelchen/api-dispatch/internal/distance"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/db"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/health"
	healthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

type Service struct {
	dataStore   *db.Store
	distanceSvc *distance.Service
}

func NewService(dataStore *db.Store, distanceSvc *distance.Service) *Service {
	return &Service{dataStore: dataStore, distanceSvc: distanceSvc}
}

func (s *Service) CreateTrips(ctx context.Context, r *v1beta1.CreateTripsRequest) (*v1beta1.CreateTripsResponse, error) {
	return s.dataStore.CreateTrips(ctx, r)
}

func (s *Service) UpdateDriverLocations(ctx context.Context, r *v1beta1.UpdateDriverLocationsRequest) (*v1beta1.UpdateDriverLocationsResponse, error) {
	return s.dataStore.UpdateDriverLocations(ctx, r)
}

func (s *Service) GetNearestDrivers(ctx context.Context, req *v1beta1.GetNearestDriversRequest) (*v1beta1.GetNearestDriversResponse, error) {
	return getNearestDrivers(ctx, req,
		s.dataStore.GetNearbyDriverLocations,
		s.distanceSvc.BetweenPoints,
		s.distanceSvc != nil,
	)
}

func (s *Service) GetNearestTrips(ctx context.Context, req *v1beta1.GetNearestTripsRequest) (*v1beta1.GetNearestTripsResponse, error) {
	return getNearestTrips(ctx, req,
		s.dataStore.GetNearbyTrips,
		s.distanceSvc.BetweenPoints,
		s.distanceSvc != nil,
	)
}

func (s *Service) Check(ctx context.Context, in *healthV1.HealthCheckRequest) (*healthV1.HealthCheckResponse, error) {
	return health.Check(ctx, in)
}

func (s *Service) Watch(in *healthV1.HealthCheckRequest, srv healthV1.Health_WatchServer) error {
	return health.Watch(in, srv)
}
