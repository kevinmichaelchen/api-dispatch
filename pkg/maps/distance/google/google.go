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

	// Batch reverse-geocode all locations
	geocoder := google.NewGeocoder(client)
	parallelizationFactor := 10
	geocodeOut, err := geocode.BatchReverseGeocode(
		ctx,
		geocoder,
		append(in.Origins, in.Destinations...),
		parallelizationFactor)
	if err != nil {
		return nil, err
	}

	var origins []string
	var destinations []string
	for idx, e := range geocodeOut {
		if idx < len(in.Origins) {
			origins = append(origins, e.PlaceID)
		} else {
			destinations = append(destinations, e.PlaceID)
		}
	}

	res, err := BetweenPlaces(ctx, client, distance.BetweenPlacesInput{
		Origins:      origins,
		Destinations: destinations,
	})
	if err != nil {
		return nil, err
	}

	return toRes(res, geocodeOut[:len(in.Origins)], geocodeOut[len(in.Origins):]), nil
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

func toRes(
	res *maps.DistanceMatrixResponse,
	originsOut []*geocode.ReverseGeocodeOutput,
	destinationsOut []*geocode.ReverseGeocodeOutput,
) *distance.MatrixResponse {
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
	var originAddresses []string
	for idx := range originsOut {
		geoResults := originsOut[idx]
		originAddresses = append(originAddresses, geoResults.FormattedAddress)
	}
	var destinationAddresses []string
	for idx := range destinationsOut {
		geoResults := destinationsOut[idx]
		destinationAddresses = append(destinationAddresses, geoResults.FormattedAddress)
	}
	return &distance.MatrixResponse{
		OriginAddresses:      originAddresses,
		DestinationAddresses: destinationAddresses,
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
