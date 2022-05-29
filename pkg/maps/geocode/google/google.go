package google

import (
	"context"
	"errors"
	maps "github.com/kevinmichaelchen/api-dispatch/pkg/maps"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/geocode"
	gMaps "googlemaps.github.io/maps"
)

var errNoResults = errors.New("no results found for coordinates")

type Geocoder struct {
	client *gMaps.Client
}

func NewGeocoder(client *gMaps.Client) *Geocoder {
	return &Geocoder{client: client}
}

func (g *Geocoder) ReverseGeocode(ctx context.Context, location maps.LatLng) (*geocode.ReverseGeocodeOutput, error) {
	results, err := g.client.ReverseGeocode(ctx, &gMaps.GeocodingRequest{
		LatLng: &gMaps.LatLng{
			Lat: location.Lat,
			Lng: location.Lng,
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
	return toReverseGeocodeOutput(results), err
}

func toReverseGeocodeOutput(in []gMaps.GeocodingResult) *geocode.ReverseGeocodeOutput {
	return &geocode.ReverseGeocodeOutput{
		// Assuming Google Maps API returns places ordered in such a way that
		// the first element is the most salient/relevant.
		PlaceID: in[0].PlaceID,
	}
}
