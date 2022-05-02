package distance

import (
	"context"
	"errors"
	"googlemaps.github.io/maps"
)

var errNoResults = errors.New("no results found for coordinates")

type Place struct {
	ID      string
	Address string
	Types   []string
}

func reverseGeocode(ctx context.Context, c *maps.Client, lat, lng float64) ([]maps.GeocodingResult, error) {
	results, err := c.ReverseGeocode(ctx, &maps.GeocodingRequest{
		Address:    "",
		Components: nil,
		Bounds:     nil,
		Region:     "",
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lng,
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
