package distance

import (
	"context"
	"errors"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"googlemaps.github.io/maps"
)

var errNoResults = errors.New("no results found for coordinates")

type Place struct {
	ID      string
	Address string
	Types   []string
}

func getFirstPlaceID(ctx context.Context, c *maps.Client, location *v1beta1.LatLng) (string, error) {
	res, err := reverseGeocode(ctx, c, location)
	if err != nil {
		return "", err
	}
	return res[0].PlaceID, nil
}

func reverseGeocode(ctx context.Context, c *maps.Client, location *v1beta1.LatLng) ([]maps.GeocodingResult, error) {
	results, err := c.ReverseGeocode(ctx, &maps.GeocodingRequest{
		Address:    "",
		Components: nil,
		Bounds:     nil,
		Region:     "",
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
