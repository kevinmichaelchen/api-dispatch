package google

import (
	"context"
	"errors"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"googlemaps.github.io/maps"
)

var errNoResults = errors.New("no results found for coordinates")

func ReverseGeocode(ctx context.Context, c *maps.Client, location *v1beta1.LatLng) ([]maps.GeocodingResult, error) {
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
