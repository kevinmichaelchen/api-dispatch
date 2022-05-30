package osrm

import (
	"context"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/geocode"
	"net/http"
)

type Geocoder struct {
	client *http.Client
}

func NewGeocoder(client *http.Client) *Geocoder {
	return &Geocoder{client: client}
}

func (g *Geocoder) ReverseGeocode(ctx context.Context, location maps.LatLng) (*geocode.ReverseGeocodeOutput, error) {
	geocoder := openstreetmap.Geocoder()

	address, err := geocoder.ReverseGeocode(location.Lat, location.Lng)
	if err != nil {
		return nil, err
	}

	return &geocode.ReverseGeocodeOutput{
		Address: *address,
	}, nil
}
