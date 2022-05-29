package distance

import (
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps"
	"time"
)

type BetweenPointsInput struct {
	Destinations []maps.LatLng
	Origins      []maps.LatLng
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

type BetweenPlacesInput struct {
	Origins      []string
	Destinations []string
}

// DistanceMatrixResponse represents a Distance Matrix API response.
type DistanceMatrixResponse struct {

	// OriginAddresses contains an array of addresses as returned by the API from
	// your original request.
	OriginAddresses []string `json:"origin_addresses"`
	// DestinationAddresses contains an array of addresses as returned by the API
	// from your original request.
	DestinationAddresses []string `json:"destination_addresses"`
	// Rows contains an array of elements.
	Rows []DistanceMatrixElementsRow `json:"rows"`
}

// DistanceMatrixElementsRow is a row of distance elements.
type DistanceMatrixElementsRow struct {
	Elements []*DistanceMatrixElement `json:"elements"`
}

// DistanceMatrixElement is the travel distance and time for a pair of origin
// and destination.
type DistanceMatrixElement struct {
	Status string `json:"status"`
	// Duration is the length of time it takes to travel this route.
	Duration time.Duration `json:"duration"`
	// DurationInTraffic is the length of time it takes to travel this route
	// considering traffic.
	DurationInTraffic time.Duration `json:"duration_in_traffic"`
	// Distance is the total distance of this route.
	Distance Distance `json:"distance"`
}
