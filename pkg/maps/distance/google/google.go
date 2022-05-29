package google

import (
	"context"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/geocode"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/geocode/google"
	"googlemaps.github.io/maps"
)

func BetweenPoints(
	ctx context.Context,
	client *maps.Client,
	in distance.BetweenPointsInput) (*distance.BetweenPointsOutput, error) {

	// Batch reverse-geocode the origins
	geocoder := google.NewGeocoder(client)

	parallelizationFactor := 10

	originsOut, err := geocode.BatchReverseGeocode(ctx, geocoder, in.Origins, parallelizationFactor)
	if err != nil {
		return nil, err
	}

	destinationsOut, err := geocode.BatchReverseGeocode(ctx, geocoder, in.Destinations, parallelizationFactor)
	if err != nil {
		return nil, err
	}

	var origins []string
	for _, e := range originsOut {
		origins = append(origins, e.PlaceID)
	}

	var destinations []string
	for _, e := range destinationsOut {
		destinations = append(destinations, e.PlaceID)
	}

	res, err := BetweenPlaces(ctx, client, distance.BetweenPlacesInput{
		Origins:      origins,
		Destinations: destinations,
	})
	if err != nil {
		return nil, err
	}

	var out []distance.Info
	for i, fromOrigin := range res.Rows {
		for j, toDestination := range fromOrigin.Elements {
			out = append(out, distance.Info{
				DistanceMeters:     toDestination.Distance.Meters,
				Duration:           toDestination.Duration,
				OriginAddress:      res.OriginAddresses[i],
				DestinationAddress: res.DestinationAddresses[j],
			})
		}
	}

	return &distance.BetweenPointsOutput{
		Info: out,
	}, nil
}

func BetweenPlaces(ctx context.Context, c *maps.Client, in distance.BetweenPlacesInput) (*maps.DistanceMatrixResponse, error) {
	var origins []string
	for _, placeID := range in.Origins {
		origins = append(origins, "place_id:"+placeID)
	}
	var destinations []string
	for _, placeID := range in.Destinations {
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
