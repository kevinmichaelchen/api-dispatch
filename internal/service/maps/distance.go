package maps

import (
	"context"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"googlemaps.github.io/maps"
	"time"
)

type Service struct {
	client *maps.Client
}

func NewService(client *maps.Client) *Service {
	return &Service{
		client: client,
	}
}

type BetweenPointsInput struct {
	Destinations []*v1beta1.LatLng
	Origins      []*v1beta1.LatLng
}

type BetweenPointsOutput struct {
	Info []Info
}

type Info struct {
	DistanceMeters     int
	Duration           time.Duration
	OriginAddress      string
	DestinationAddress string
}

func (s *Service) BetweenPoints(ctx context.Context, in BetweenPointsInput) (*BetweenPointsOutput, error) {
	origins, err := locationsToPlaceIDs(ctx, s.client, in.Origins)
	if err != nil {
		return nil, err
	}

	destinations, err := locationsToPlaceIDs(ctx, s.client, in.Destinations)
	if err != nil {
		return nil, err
	}

	// TODO if len(origins) > 25 || len(destinations) > 25, we need to partition/batch

	res, err := betweenPlaces(ctx, s.client, betweenPlacesInput{
		originPlaceIDs:      origins,
		destinationPlaceIDs: destinations,
	})
	if err != nil {
		return nil, err
	}

	var out []Info
	for i, fromOrigin := range res.Rows {
		for j, toDestination := range fromOrigin.Elements {
			out = append(out, Info{
				DistanceMeters:     toDestination.Distance.Meters,
				Duration:           toDestination.Duration,
				OriginAddress:      res.OriginAddresses[i],
				DestinationAddress: res.DestinationAddresses[j],
			})
		}
	}

	return &BetweenPointsOutput{
		Info: out,
	}, nil
}

type betweenPlacesInput struct {
	originPlaceIDs      []string
	destinationPlaceIDs []string
}

func betweenPlaces(ctx context.Context, c *maps.Client, in betweenPlacesInput) (*maps.DistanceMatrixResponse, error) {
	var origins []string
	for _, placeID := range in.originPlaceIDs {
		origins = append(origins, "place_id:"+placeID)
	}
	var destinations []string
	for _, placeID := range in.destinationPlaceIDs {
		destinations = append(destinations, "place_id:"+placeID)
	}
	return c.DistanceMatrix(ctx, &maps.DistanceMatrixRequest{
		Origins:                  origins,
		Destinations:             destinations,
		Mode:                     "",
		Language:                 "",
		Avoid:                    "",
		Units:                    "",
		DepartureTime:            "",
		ArrivalTime:              "",
		TrafficModel:             "",
		TransitMode:              nil,
		TransitRoutingPreference: "",
	})
}
