package service

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	errUnsupportedKValue = errors.New("unsupported k-value")
)

type getNearbyDriverLocationsOutput struct {
	r7k1Cells  models.DriverLocationSlice
	r8k1Cells  models.DriverLocationSlice
	r8k2Cells  models.DriverLocationSlice
	r9k1Cells  models.DriverLocationSlice
	r9k2Cells  models.DriverLocationSlice
	r10k1Cells models.DriverLocationSlice
	r10k2Cells models.DriverLocationSlice
}

func (s *Service) getNearbyDriverLocations(ctx context.Context, location *v1beta1.LatLng) (*getNearbyDriverLocationsOutput, error) {
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
	return &getNearbyDriverLocationsOutput{
		r7k1Cells:  r7k1Cells,
		r8k1Cells:  r8k1Cells,
		r8k2Cells:  r8k2Cells,
		r9k1Cells:  r9k1Cells,
		r9k2Cells:  r9k2Cells,
		r10k1Cells: r10k1Cells,
		r10k2Cells: r10k2Cells,
	}, nil
}

func (s *Service) getNearbyDriverLocationsHelper(ctx context.Context, l *v1beta1.LatLng, res int, k int) (models.DriverLocationSlice, error) {
	if k < 1 || k > 2 {
		return nil, errUnsupportedKValue
	}
	cell := getCell(l, res)
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

type mergeDriversInput struct {
	drivers models.DriverLocationSlice
	res     int
	kValue  int
}

func mergeDrivers(location *v1beta1.LatLng, in ...mergeDriversInput) []*v1beta1.SearchResult {
	cache := make(map[string]*v1beta1.SearchResult)
	for _, mi := range in {
		for _, dl := range mi.drivers {
			extant, exists := cache[dl.DriverID]
			// if we've already recorded the driver appearing in a
			// higher-resolution neighbor, skip
			if exists {
				// prefer higher-res immediate neighbors
				if extant.GetResolution() > int32(mi.res) {
					continue
				}
				// if they're the same res, prefer those in 1-ring over the
				// 2-ring
				if extant.GetResolution() == int32(mi.res) {
					if extant.GetKValue() >= int32(mi.kValue) {
						continue
					}
				}
			}
			latLng := &v1beta1.LatLng{
				Latitude:  dl.Latitude,
				Longitude: dl.Longitude,
			}
			cache[dl.DriverID] = &v1beta1.SearchResult{
				Payload: &v1beta1.SearchResult_Driver{
					Driver: &v1beta1.DriverLocation{
						DriverId:            dl.DriverID,
						MostRecentHeartbeat: timestamppb.New(dl.CreatedAt),
						CurrentLocation:     latLng,
					},
				},
				DistanceMeters: pointDistance(location, latLng),
				Location:       latLng,
				Resolution:     int32(mi.res),
				KValue:         int32(mi.kValue),
			}
		}
	}
	var out []*v1beta1.SearchResult
	for _, e := range cache {
		out = append(out, e)
	}
	return out
}

type getNearbyTripsOutput struct {
	r7k1Cells  models.TripSlice
	r8k1Cells  models.TripSlice
	r8k2Cells  models.TripSlice
	r9k1Cells  models.TripSlice
	r9k2Cells  models.TripSlice
	r10k1Cells models.TripSlice
	r10k2Cells models.TripSlice
}

func (s *Service) getNearbyTrips(ctx context.Context, location *v1beta1.LatLng) (*getNearbyTripsOutput, error) {
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
	return &getNearbyTripsOutput{
		r7k1Cells:  r7k1Cells,
		r8k1Cells:  r8k1Cells,
		r8k2Cells:  r8k2Cells,
		r9k1Cells:  r9k1Cells,
		r9k2Cells:  r9k2Cells,
		r10k1Cells: r10k1Cells,
		r10k2Cells: r10k2Cells,
	}, nil
}

func (s *Service) getNearbyTripsHelper(ctx context.Context, l *v1beta1.LatLng, res int, k int) (models.TripSlice, error) {
	if k < 1 || k > 2 {
		return nil, errUnsupportedKValue
	}
	cell := getCell(l, res)
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

type mergeTripsInput struct {
	trips  models.TripSlice
	res    int
	kValue int
}

func mergeTrips(location *v1beta1.LatLng, in ...mergeTripsInput) []*v1beta1.SearchResult {
	cache := make(map[string]*v1beta1.SearchResult)
	for _, mi := range in {
		for _, e := range mi.trips {
			extant, exists := cache[e.ID]
			// if we've already recorded the driver appearing in a
			// higher-resolution neighbor, skip
			if exists {
				// prefer higher-res immediate neighbors
				if extant.GetResolution() > int32(mi.res) {
					continue
				}
				// if they're the same res, prefer those in 1-ring over the
				// 2-ring
				if extant.GetResolution() == int32(mi.res) {
					if extant.GetKValue() >= int32(mi.kValue) {
						continue
					}
				}
			}
			latLng := &v1beta1.LatLng{
				Latitude:  e.Latitude,
				Longitude: e.Longitude,
			}
			cache[e.ID] = &v1beta1.SearchResult{
				Payload: &v1beta1.SearchResult_Trip{
					Trip: &v1beta1.Trip{
						Id:              e.ID,
						PickupLocation:  latLng,
						ScheduledFor:    timestamppb.New(e.ScheduledFor),
						ExpectedPayment: e.ExpectedPay,
					},
				},
				DistanceMeters: pointDistance(location, latLng),
				Location:       latLng,
				Resolution:     int32(mi.res),
				KValue:         int32(mi.kValue),
			}
		}
	}
	var out []*v1beta1.SearchResult
	for _, e := range cache {
		out = append(out, e)
	}
	return out
}
