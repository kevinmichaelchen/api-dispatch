package service

import (
	"context"
	"fmt"
	"github.com/kevinmichaelchen/api-dispatch/internal/distance"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"sort"
)

const (
	maxResults = 100
)

func (s *Service) GetNearestDrivers(ctx context.Context, req *v1beta1.GetNearestDriversRequest) (*v1beta1.GetNearestDriversResponse, error) {
	err := validateGetNearestDriversRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// TODO parallelize with errgroup
	r7k1Cells, err := s.getNearbyDriverLocations(ctx, req.GetLocation(), 7, 1)
	if err != nil {
		return nil, err
	}
	r8k1Cells, err := s.getNearbyDriverLocations(ctx, req.GetLocation(), 8, 1)
	if err != nil {
		return nil, err
	}
	r8k2Cells, err := s.getNearbyDriverLocations(ctx, req.GetLocation(), 8, 2)
	if err != nil {
		return nil, err
	}
	r9k1Cells, err := s.getNearbyDriverLocations(ctx, req.GetLocation(), 9, 1)
	if err != nil {
		return nil, err
	}
	r9k2Cells, err := s.getNearbyDriverLocations(ctx, req.GetLocation(), 9, 2)
	if err != nil {
		return nil, err
	}
	r10k1Cells, err := s.getNearbyDriverLocations(ctx, req.GetLocation(), 10, 1)
	if err != nil {
		return nil, err
	}
	r10k2Cells, err := s.getNearbyDriverLocations(ctx, req.GetLocation(), 10, 2)
	if err != nil {
		return nil, err
	}

	results := merge(
		req,
		mergeInput{drivers: r7k1Cells, res: 7, kValue: 1},
		mergeInput{drivers: r8k1Cells, res: 8, kValue: 1},
		mergeInput{drivers: r8k2Cells, res: 8, kValue: 2},
		mergeInput{drivers: r9k1Cells, res: 9, kValue: 1},
		mergeInput{drivers: r9k2Cells, res: 9, kValue: 2},
		mergeInput{drivers: r10k1Cells, res: 10, kValue: 1},
		mergeInput{drivers: r10k2Cells, res: 10, kValue: 2},
	)

	// Apply server-side results limit
	if len(results) > int(maxResults) {
		results = results[:maxResults]
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(results, func(i, j int) bool {
		a := results[i]
		b := results[j]
		if a.Resolution == b.Resolution {
			return a.KValue > b.KValue
		}
		return a.Resolution > b.Resolution
	})

	// Enrich results with distance/duration info from Google Maps API
	var driverLocations []*v1beta1.LatLng
	for _, result := range results {
		driverLocations = append(driverLocations, result.GetDriverLocation())
	}
	var pickupAddress string
	if s.distanceSvc != nil {
		out, err := s.distanceSvc.BetweenPoints(ctx, distance.BetweenPointsInput{
			PickupLocation:  req.GetLocation(),
			DriverLocations: driverLocations,
		})
		if err != nil {
			return nil, err
		}
		for i, info := range out.Info {
			results[i].Duration = durationpb.New(info.Duration)
			results[i].DistanceMeters = float64(info.DistanceMeters)
			results[i].Address = info.OriginAddress
		}
		pickupAddress = out.DestinationAddress
	}

	// Re-sort by duration
	sort.SliceStable(results, func(i, j int) bool {
		a := results[i]
		b := results[j]
		return a.Duration.AsDuration() < b.Duration.AsDuration()
	})

	// Apply client-side limits
	// TODO do not let client exceed server-side max limit
	if len(results) > int(req.GetLimit()) {
		results = results[:req.GetLimit()]
	}

	return &v1beta1.GetNearestDriversResponse{
		Results:       results,
		PickupAddress: pickupAddress,
	}, nil
}

type mergeInput struct {
	drivers models.DriverLocationSlice
	res     int
	kValue  int
}

func merge(r *v1beta1.GetNearestDriversRequest, in ...mergeInput) []*v1beta1.SearchResult {
	cache := make(map[string]*v1beta1.SearchResult)
	for _, mi := range in {
		for _, dl := range mi.drivers {
			extant, exists := cache[dl.DriverID]
			// if we've already recorded the driver appearing in a
			// higher-resolution neighbor, skip
			if exists {
				// prefer higher-res immediate neighbors
				if extant.Resolution > int32(mi.res) {
					continue
				}
				// if they're the same res, prefer those in 1-ring over the
				// 2-ring
				if extant.Resolution == int32(mi.res) {
					if extant.KValue >= int32(mi.kValue) {
						continue
					}
				}
			}
			latLng := &v1beta1.LatLng{
				Latitude:  dl.Latitude,
				Longitude: dl.Longitude,
			}
			cache[dl.DriverID] = &v1beta1.SearchResult{
				DriverId:       dl.DriverID,
				DistanceMeters: pointDistance(r.GetLocation(), latLng),
				DriverLocation: latLng,
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

// k should be either 1 or 2
func (s *Service) getNearbyDriverLocations(ctx context.Context, l *v1beta1.LatLng, res int, k int) (models.DriverLocationSlice, error) {
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

type DriverID string

func toSearchResults(
	r *v1beta1.GetNearestDriversRequest,
	in models.DriverLocationSlice,
	res int,
	k int) map[DriverID]*v1beta1.SearchResult {
	out := make(map[DriverID]*v1beta1.SearchResult)
	for _, e := range in {
		driverLocation := &v1beta1.LatLng{
			Latitude:  e.Latitude,
			Longitude: e.Longitude,
		}
		out[DriverID(e.DriverID)] = &v1beta1.SearchResult{
			DriverId:       e.DriverID,
			DistanceMeters: pointDistance(r.GetLocation(), driverLocation),
			DriverLocation: driverLocation,
			Resolution:     int32(res),
			KValue:         int32(k),
		}
	}
	return out
}
