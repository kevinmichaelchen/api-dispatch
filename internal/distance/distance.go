package distance

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
	PickupLocations []*v1beta1.LatLng
	DriverLocations []*v1beta1.LatLng
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
	driverPlaceIDs, err := locationsToPlaceIDs(ctx, s.client, in.DriverLocations)
	if err != nil {
		return nil, err
	}

	pickupPlaceIDs, err := locationsToPlaceIDs(ctx, s.client, in.PickupLocations)
	if err != nil {
		return nil, err
	}

	res, err := betweenPlaces(ctx, s.client, betweenPlacesInput{
		originPlaceIDs:      driverPlaceIDs,
		destinationPlaceIDs: pickupPlaceIDs,
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
