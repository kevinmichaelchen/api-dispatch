package geo

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance/google"
	gMaps "googlemaps.github.io/maps"
)

type Service struct {
	client *gMaps.Client
}

func NewService(client *gMaps.Client) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) BetweenPoints(ctx context.Context, in distance.BetweenPointsInput) (*distance.BetweenPointsOutput, error) {
	err := validate(in)
	if err != nil {
		return nil, err
	}

	// TODO if len(origins) > 25 || len(destinations) > 25, we need to partition/batch

	return google.BetweenPoints(ctx, s.client, in)
}

func validate(i distance.BetweenPointsInput) error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Destinations, validation.Required, validation.Length(1, 0)),
		validation.Field(&i.Origins, validation.Required, validation.Length(1, 0)),
	)
}
