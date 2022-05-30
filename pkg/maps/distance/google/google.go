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
	in distance.BetweenPointsInput) (*distance.MatrixResponse, error) {

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

	return toRes(res), nil

	// TODO throw in some reverse-geocoding for origins+destination addresses
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

func toRes(res *maps.DistanceMatrixResponse) *distance.MatrixResponse {
	var rows []distance.MatrixElementsRow
	for i := range res.Rows {
		row := res.Rows[i]
		var elements []distance.MatrixElement
		for j := range row.Elements {
			elem := row.Elements[j]
			elements = append(elements, toElem(elem))
		}
		rows = append(rows, distance.MatrixElementsRow{Elements: elements})
	}
	return &distance.MatrixResponse{
		OriginAddresses:      nil,
		DestinationAddresses: nil,
		Rows:                 rows,
	}
}

func toElem(res *maps.DistanceMatrixElement) distance.MatrixElement {
	return distance.MatrixElement{
		Status:            res.Status,
		Duration:          res.Duration,
		DurationInTraffic: res.DurationInTraffic,
		Distance:          res.Distance.Meters,
	}
}
