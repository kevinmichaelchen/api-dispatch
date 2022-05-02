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

type BetweenPointsOutput struct {
	DistanceMeters int
	Duration       time.Duration
}

func (s *Service) BetweenPoints(ctx context.Context, p1, p2 *v1beta1.LatLng) (*BetweenPointsOutput, error) {
	place1Res, err := reverseGeocode(ctx, s.client, p1.GetLatitude(), p1.GetLongitude())
	if err != nil {
		return nil, err
	}
	place2Res, err := reverseGeocode(ctx, s.client, p2.GetLatitude(), p2.GetLongitude())
	if err != nil {
		return nil, err
	}
	place1 := place1Res[0].PlaceID
	place2 := place2Res[0].PlaceID

	res, err := betweenPlaces(ctx, s.client, place1, place2)
	if err != nil {
		return nil, err
	}
	elem := res.Rows[0].Elements[0]
	return &BetweenPointsOutput{
		DistanceMeters: elem.Distance.Meters,
		Duration:       elem.Duration,
	}, nil
}

func betweenPlaces(ctx context.Context, c *maps.Client, placeID1, placeID2 string) (*maps.DistanceMatrixResponse, error) {
	return c.DistanceMatrix(ctx, &maps.DistanceMatrixRequest{
		// https://developers.google.com/maps/documentation/distance-matrix/distance-matrix#origins
		Origins: []string{
			"place_id:" + placeID1,
		},
		Destinations: []string{
			"place_id:" + placeID2,
		},
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
