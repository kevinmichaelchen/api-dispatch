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
	PickupLocation  *v1beta1.LatLng
	DriverLocations []*v1beta1.LatLng
}

type BetweenPointsOutput struct {
	Info               []Info
	DestinationAddress string
}

type Info struct {
	DistanceMeters int
	Duration       time.Duration
	OriginAddress  string
}

func (s *Service) BetweenPoints(ctx context.Context, in BetweenPointsInput) (*BetweenPointsOutput, error) {
	var driverPlaceIDs []string
	for _, dl := range in.DriverLocations {
		driverPlaceID, err := getFirstPlaceID(ctx, s.client, dl)
		if err != nil {
			return nil, err
		}
		driverPlaceIDs = append(driverPlaceIDs, driverPlaceID)
	}

	pickupPlaceID, err := getFirstPlaceID(ctx, s.client, in.PickupLocation)
	if err != nil {
		return nil, err
	}

	res, err := betweenPlaces(ctx, s.client, betweenPlacesInput{
		originPlaceIDs:     driverPlaceIDs,
		destinationPlaceID: pickupPlaceID,
	})
	if err != nil {
		return nil, err
	}

	var out []Info
	for i, fromOrigin := range res.Rows {
		for _, toDestination := range fromOrigin.Elements {
			out = append(out, Info{
				DistanceMeters: toDestination.Distance.Meters,
				Duration:       toDestination.Duration,
				OriginAddress:  res.OriginAddresses[i],
			})
		}
	}

	return &BetweenPointsOutput{
		DestinationAddress: res.DestinationAddresses[0],
		Info:               out,
	}, nil
}

type betweenPlacesInput struct {
	originPlaceIDs     []string
	destinationPlaceID string
}

func betweenPlaces(ctx context.Context, c *maps.Client, in betweenPlacesInput) (*maps.DistanceMatrixResponse, error) {
	var origins []string
	for _, placeID := range in.originPlaceIDs {
		origins = append(origins, "place_id:"+placeID)
	}
	destinations := []string{"place_id:" + in.destinationPlaceID}
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
