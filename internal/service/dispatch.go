package service

import (
	"context"
	"fmt"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"sort"
)

func (s *Service) Dispatch(ctx context.Context, r *v1beta1.DispatchRequest) (*v1beta1.DispatchResponse, error) {
	// TODO parallelize with errgroup
	r7Neighbors, err := s.getNearbyDriverLocations(ctx, r.GetLocation(), 7)
	if err != nil {
		return nil, err
	}
	r8Neighbors, err := s.getNearbyDriverLocations(ctx, r.GetLocation(), 8)
	if err != nil {
		return nil, err
	}
	r9Neighbors, err := s.getNearbyDriverLocations(ctx, r.GetLocation(), 9)
	if err != nil {
		return nil, err
	}
	r10Neighbors, err := s.getNearbyDriverLocations(ctx, r.GetLocation(), 10)
	if err != nil {
		return nil, err
	}

	results := merge(
		r,
		mergeInput{drivers: r7Neighbors, res: 7},
		mergeInput{drivers: r8Neighbors, res: 8},
		mergeInput{drivers: r9Neighbors, res: 9},
		mergeInput{drivers: r10Neighbors, res: 10},
	)

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(results, func(i, j int) bool {
		return results[i].Resolution > results[j].Resolution
	})

	return &v1beta1.DispatchResponse{
		Results: results,
	}, nil
}

type mergeInput struct {
	drivers models.DriverLocationSlice
	res     int
}

func merge(r *v1beta1.DispatchRequest, in ...mergeInput) []*v1beta1.SearchResult {
	cache := make(map[string]*v1beta1.SearchResult)
	for _, mi := range in {
		for _, dl := range mi.drivers {
			extant, exists := cache[dl.DriverID]
			// if we've already recorded the driver appearing in a
			// higher-resolution neighbor, skip
			if exists && extant.Resolution >= int32(mi.res) {
				continue
			}
			cache[dl.DriverID] = &v1beta1.SearchResult{
				DriverId: dl.DriverID,
				DistanceMiles: distance(r.GetLocation(), &v1beta1.LatLng{
					Latitude:  dl.Latitude,
					Longitude: dl.Longitude,
				}),
				Resolution: int32(mi.res),
			}
		}
	}
	var out []*v1beta1.SearchResult
	for _, e := range cache {
		out = append(out, e)
	}
	return out
}

func (s *Service) getNearbyDriverLocations(ctx context.Context, l *v1beta1.LatLng, res int) (models.DriverLocationSlice, error) {
	cell := getCell(l, res)
	var obj models.DriverLocationSlice
	queryTemplate := `
SELECT driver_id, latitude, longitude
FROM driver_location 
WHERE 
  $1 = ANY (r%d_k1_neighbors)
`
	query := fmt.Sprintf(queryTemplate, res)
	err := queries.Raw(query, cell).Bind(ctx, s.db, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type DriverID string

func toSearchResults(r *v1beta1.DispatchRequest, in models.DriverLocationSlice, res int) map[DriverID]*v1beta1.SearchResult {
	out := make(map[DriverID]*v1beta1.SearchResult)
	for _, e := range in {
		out[DriverID(e.DriverID)] = &v1beta1.SearchResult{
			DriverId: e.DriverID,
			DistanceMiles: distance(r.GetLocation(), &v1beta1.LatLng{
				Latitude:  e.Latitude,
				Longitude: e.Longitude,
			}),
			Resolution: int32(res),
		}
	}
	return out
}
