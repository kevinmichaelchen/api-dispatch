package google

import (
	"context"
	"errors"
	"github.com/codingsince1985/geo-golang"
	maps "github.com/kevinmichaelchen/api-dispatch/pkg/maps"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/geocode"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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
	tr := otel.Tracer("")
	ctx, span := tr.Start(ctx, "ReverseGeocode")
	span.SetAttributes(
		attribute.Key("lat").Float64(location.Lat),
		attribute.Key("lon").Float64(location.Lng),
	)
	defer span.End()

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
	// Assuming Google Maps API returns places ordered in such a way that
	// the first element is the most salient/relevant.
	bestResult := in[0]
	return &geocode.ReverseGeocodeOutput{
		PlaceID: bestResult.PlaceID,
		Address: geo.Address{
			FormattedAddress: bestResult.FormattedAddress,
			// TODO set other components
			Street:        "",
			HouseNumber:   "",
			Suburb:        "",
			Postcode:      "",
			State:         "",
			StateCode:     "",
			StateDistrict: "",
			County:        "",
			Country:       "",
			CountryCode:   "",
			City:          "",
		},
	}
}
