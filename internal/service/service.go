package service

import (
	"context"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/db"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/geo"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/health"
	"google.golang.org/grpc/codes"
	healthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Service struct {
	dataStore   *db.Store
	distanceSvc *geo.Service
}

func NewService(dataStore *db.Store, distanceSvc *geo.Service) *Service {
	return &Service{dataStore: dataStore, distanceSvc: distanceSvc}
}

type Validater interface {
	Validate() error
}

func validate(m proto.Message, r Validater) error {
	//name := m.ProtoReflect().Type().Descriptor().Name()
	err := r.Validate()
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}

func (s *Service) CreateTrips(ctx context.Context, r *v1beta1.CreateTripsRequest) (*v1beta1.CreateTripsResponse, error) {
	err := validate(r, r)
	if err != nil {
		return nil, err
	}
	return s.dataStore.CreateTrips(ctx, r)
}

func (s *Service) UpdateDriverLocations(ctx context.Context, r *v1beta1.UpdateDriverLocationsRequest) (*v1beta1.UpdateDriverLocationsResponse, error) {
	err := validate(r, r)
	if err != nil {
		return nil, err
	}

	return s.dataStore.UpdateDriverLocations(ctx, r)
}

func (s *Service) Check(ctx context.Context, in *healthV1.HealthCheckRequest) (*healthV1.HealthCheckResponse, error) {
	return health.Check(ctx, in)
}

func (s *Service) Watch(in *healthV1.HealthCheckRequest, srv healthV1.Health_WatchServer) error {
	return health.Watch(in, srv)
}
