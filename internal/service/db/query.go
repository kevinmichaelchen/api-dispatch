package db

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/h3"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"go.opentelemetry.io/otel"
)

var (
	errUnsupportedKValue = errors.New("unsupported k-value")
)

type GetNearbyDriverLocationsOutput struct {
	// Drivers in the passenger's Resolution-7 k=1 k-ring
	R7K1 models.DriverLocationSlice
	// Drivers in the passenger's Resolution-8 k=1 k-ring
	R8K1 models.DriverLocationSlice
	// Drivers in the passenger's Resolution-8 k=2 k-ring
	R8K2 models.DriverLocationSlice
	// Drivers in the passenger's Resolution-9 k=1 k-ring
	R9K1 models.DriverLocationSlice
	// Drivers in the passenger's Resolution-9 k=2 k-ring
	R9K2 models.DriverLocationSlice
	// Drivers in the passenger's Resolution-10 k=1 k-ring
	R10K1 models.DriverLocationSlice
	// Drivers in the passenger's Resolution-10 k=2 k-ring
	R10K2 models.DriverLocationSlice
}

func (s *Store) GetNearbyDriverLocations(ctx context.Context, location *v1beta1.LatLng) (*GetNearbyDriverLocationsOutput, error) {
	tr := otel.Tracer("")
	ctx, span := tr.Start(ctx, "GetNearbyDriverLocations")
	defer span.End()

	// TODO parallelize with errgroup
	r7k1Cells, err := s.getNearbyDriverLocationsHelper(ctx, location, 7, 1)
	if err != nil {
		return nil, err
	}
	r8k1Cells, err := s.getNearbyDriverLocationsHelper(ctx, location, 8, 1)
	if err != nil {
		return nil, err
	}
	r8k2Cells, err := s.getNearbyDriverLocationsHelper(ctx, location, 8, 2)
	if err != nil {
		return nil, err
	}
	r9k1Cells, err := s.getNearbyDriverLocationsHelper(ctx, location, 9, 1)
	if err != nil {
		return nil, err
	}
	r9k2Cells, err := s.getNearbyDriverLocationsHelper(ctx, location, 9, 2)
	if err != nil {
		return nil, err
	}
	r10k1Cells, err := s.getNearbyDriverLocationsHelper(ctx, location, 10, 1)
	if err != nil {
		return nil, err
	}
	r10k2Cells, err := s.getNearbyDriverLocationsHelper(ctx, location, 10, 2)
	if err != nil {
		return nil, err
	}
	return &GetNearbyDriverLocationsOutput{
		R7K1:  r7k1Cells,
		R8K1:  r8k1Cells,
		R8K2:  r8k2Cells,
		R9K1:  r9k1Cells,
		R9K2:  r9k2Cells,
		R10K1: r10k1Cells,
		R10K2: r10k2Cells,
	}, nil
}

func (s *Store) getNearbyDriverLocationsHelper(ctx context.Context, l *v1beta1.LatLng, res int, k int) (models.DriverLocationSlice, error) {
	// TODO filter out offline drivers or busy (currently-on-a-trip) drivers
	if k < 1 || k > 2 {
		return nil, errUnsupportedKValue
	}
	cell := h3.GetCell(l, res)
	var obj models.DriverLocationSlice
	queryTemplate := `
SELECT driver_id, latitude, longitude
FROM driver_location 
WHERE 
  $1 = ANY (r%d_k%d_neighbors)
`
	query := fmt.Sprintf(queryTemplate, res, k)
	err := queries.Raw(query, cell).Bind(ctx, s.db, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type GetNearbyTripsOutput struct {
	// Trips in the driver's Resolution-7 k=1 k-ring
	R7K1 models.TripSlice
	// Trips in the driver's Resolution-8 k=1 k-ring
	R8K1 models.TripSlice
	// Trips in the driver's Resolution-8 k=2 k-ring
	R8K2 models.TripSlice
	// Trips in the driver's Resolution-9 k=1 k-ring
	R9K1 models.TripSlice
	// Trips in the driver's Resolution-9 k=2 k-ring
	R9K2 models.TripSlice
	// Trips in the driver's Resolution-10 k=1 k-ring
	R10K1 models.TripSlice
	// Trips in the driver's Resolution-10 k=2 k-ring
	R10K2 models.TripSlice
}

func (s *Store) GetNearbyTrips(ctx context.Context, location *v1beta1.LatLng) (*GetNearbyTripsOutput, error) {
	tr := otel.Tracer("")
	ctx, span := tr.Start(ctx, "GetNearbyTrips")
	defer span.End()

	// TODO parallelize with errgroup
	r7k1Cells, err := s.getNearbyTripsHelper(ctx, location, 7, 1)
	if err != nil {
		return nil, err
	}
	r8k1Cells, err := s.getNearbyTripsHelper(ctx, location, 8, 1)
	if err != nil {
		return nil, err
	}
	r8k2Cells, err := s.getNearbyTripsHelper(ctx, location, 8, 2)
	if err != nil {
		return nil, err
	}
	r9k1Cells, err := s.getNearbyTripsHelper(ctx, location, 9, 1)
	if err != nil {
		return nil, err
	}
	r9k2Cells, err := s.getNearbyTripsHelper(ctx, location, 9, 2)
	if err != nil {
		return nil, err
	}
	r10k1Cells, err := s.getNearbyTripsHelper(ctx, location, 10, 1)
	if err != nil {
		return nil, err
	}
	r10k2Cells, err := s.getNearbyTripsHelper(ctx, location, 10, 2)
	if err != nil {
		return nil, err
	}
	return &GetNearbyTripsOutput{
		R7K1:  r7k1Cells,
		R8K1:  r8k1Cells,
		R8K2:  r8k2Cells,
		R9K1:  r9k1Cells,
		R9K2:  r9k2Cells,
		R10K1: r10k1Cells,
		R10K2: r10k2Cells,
	}, nil
}

// TODO trips that are too far in the past (or in a terminal state) should be filtered out
func (s *Store) getNearbyTripsHelper(ctx context.Context, l *v1beta1.LatLng, res int, k int) (models.TripSlice, error) {
	if k < 1 || k > 2 {
		return nil, errUnsupportedKValue
	}
	cell := h3.GetCell(l, res)
	var obj models.TripSlice
	queryTemplate := `
SELECT id, latitude, longitude, scheduled_for, expected_pay 
FROM trip 
WHERE 
  $1 = ANY (r%d_k%d_neighbors)
`
	query := fmt.Sprintf(queryTemplate, res, k)
	err := queries.Raw(query, cell).Bind(ctx, s.db, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
