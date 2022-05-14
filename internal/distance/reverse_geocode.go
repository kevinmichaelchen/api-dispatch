package distance

import (
	"context"
	"errors"
	"fmt"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"golang.org/x/sync/errgroup"
	"googlemaps.github.io/maps"
	"strconv"
	"sync"
	"sync/atomic"
)

var errNoResults = errors.New("no results found for coordinates")

type Place struct {
	ID      string
	Address string
	Types   []string
}

type LatLng struct {
	Lat float64
	Lng float64
}

// locationsToPlaceIDs reverse geocodes a list of geographic coordinates.
// It uses parallelization with the errgroup package, since Google Maps does not
// offer a way to reverse geocode in bulk, and since each individual request
// would take roughly 150 ms.
// Inspired by:
// https://www.fullstory.com/blog/why-errgroup-withcontext-in-golang-server-handlers/
func locationsToPlaceIDs(ctx context.Context, c *maps.Client, locations []*v1beta1.LatLng) ([]string, error) {
	g, ctx := errgroup.WithContext(ctx)

	locationsChan := make(chan *v1beta1.LatLng)

	// Step 1: Produce
	g.Go(func() error {
		defer close(locationsChan)
		for _, location := range locations {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case locationsChan <- location:
			}
		}
		return nil
	})

	type Result struct {
		LatLng           LatLng
		GeocodingResults []maps.GeocodingResult
	}
	results := make(chan Result)

	// Step 2: Map
	nWorkers := 5 // parallelization factor
	workers := int32(nWorkers)
	for i := 0; i < nWorkers; i++ {
		g.Go(func() error {
			defer func() {
				// Last one out closes shop
				if atomic.AddInt32(&workers, -1) == 0 {
					close(results)
				}
			}()

			for location := range locationsChan {
				geocodingResults, err := reverseGeocode(ctx, c, location)
				if err != nil {
					return fmt.Errorf("failed to reverse geocode location: %w", err)
				} else {
					result := Result{
						LatLng: LatLng{
							Lat: location.GetLatitude(),
							Lng: location.GetLongitude(),
						},
						GeocodingResults: geocodingResults,
					}
					select {
					case <-ctx.Done():
						return ctx.Err()
					case results <- result:
					}
				}
			}
			return nil
		})
	}

	// Step 3: Reduce
	// A normal Go map isn't thread-safe, so we use sync.Map
	ret := new(sync.Map)
	g.Go(func() error {
		for result := range results {
			ret.Store(result.LatLng, result.GeocodingResults)
		}
		return nil
	})

	// Wait blocks until all function calls from the Go method have returned, then
	// returns the first non-nil error (if any) from them.
	err := g.Wait()
	if err != nil {
		return nil, err
	}

	// Step 4: Convert sync.Map into a list
	// The order of outputs has to correspond with the order of inputs
	// e.g., if the input was [point1, point2] then the output should be
	// [place1ID, place2ID]
	var out []string
	for _, location := range locations {
		lat := location.GetLatitude()
		lng := location.GetLongitude()
		val, ok := ret.Load(LatLng{Lat: lat, Lng: lng})
		if !ok {
			return nil, fmt.Errorf("could not find lat/lng for (%s, %s)",
				strconv.FormatFloat(lat, 'f', -1, 64),
				strconv.FormatFloat(lng, 'f', -1, 64),
			)
		}
		geocodingResults, ok := val.([]maps.GeocodingResult)
		if !ok {
			return nil, errors.New("expected sync.Map values to be []maps.GeocodingResult")
		}
		if len(geocodingResults) == 0 {
			return nil, errors.New("sync.Map value was empty slice of []maps.GeocodingResult")
		}

		// Assuming Google Maps API returns places ordered in such a way that
		// the first element is the most salient/relevant.
		out = append(out, geocodingResults[0].PlaceID)
	}
	return out, nil
}

func reverseGeocode(ctx context.Context, c *maps.Client, location *v1beta1.LatLng) ([]maps.GeocodingResult, error) {
	results, err := c.ReverseGeocode(ctx, &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: location.GetLatitude(),
			Lng: location.GetLongitude(),
		},
		ResultType:   nil,
		LocationType: nil,
		PlaceID:      "",
		Language:     "",
		Custom:       nil,
	})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errNoResults
	}
	return results, err
}
