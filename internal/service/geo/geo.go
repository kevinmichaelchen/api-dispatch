package geo

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance/google"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance/osrm"
	gMaps "googlemaps.github.io/maps"
	"net/http"
)

type Service struct {
	googleClient *gMaps.Client
	httpClient   *http.Client
}

func NewService(client *gMaps.Client, httpClient *http.Client) *Service {
	return &Service{
		googleClient: client,
		httpClient:   httpClient,
	}
}

func (s *Service) BetweenPoints(ctx context.Context, in distance.BetweenPointsInput) (*distance.BetweenPointsOutput, error) {
	err := validate(in)
	if err != nil {
		return nil, err
	}

	// TODO if len(origins) > 25 || len(destinations) > 25, we need to partition/batch

	// Use Google Maps if there's an API key available
	if s.googleClient != nil {
		return google.BetweenPoints(ctx, s.googleClient, in)
	}

	// Otherwise we'll back to using Open Source Routing Machine (OSRM)
	return osrm.BetweenPoints(ctx, s.httpClient, in)
}

func validate(i distance.BetweenPointsInput) error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Destinations, validation.Required, validation.Length(1, 25)),
		validation.Field(&i.Origins, validation.Required, validation.Length(1, 25)),
	)
}
