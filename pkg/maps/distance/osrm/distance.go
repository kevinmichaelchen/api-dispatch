package osrm

import (
	"context"
	"errors"
	"fmt"
	osrm "github.com/gojuno/go.osrm"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance"
	geo "github.com/paulmach/go.geo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	errFailedRequest = errors.New("failed OSRM request")
)

func BetweenPoints(
	ctx context.Context,
	httpClient *http.Client,
	in distance.BetweenPointsInput) (*distance.MatrixResponse, error) {
	logger := ctxzap.Extract(ctx)

	serverURL := "https://router.project-osrm.org"
	client := osrm.NewWithConfig(osrm.Config{
		ServerURL: serverURL,
		Client:    httpClient,
	})

	res, err := client.Table(ctx, toTableReq(in))

	if err != nil {
		return nil, fmt.Errorf("failed OSRM request: %w", err)
	}

	if res.Code != "Ok" {
		logger.Error("received non-ok OSRM response",
			zap.String("error.code", res.Code),
			zap.String("error.msg", res.Message),
		)
		return nil, errFailedRequest
	}

	// TODO throw in some reverse-geocoding for origins+destination addresses

	return fromTableRes(res), nil
}

func toTableReq(in distance.BetweenPointsInput) osrm.TableRequest {
	var pointSet geo.PointSet
	for _, p := range in.Origins {
		pointSet = append(pointSet, geo.Point{p.Lng, p.Lat})
	}
	for _, p := range in.Destinations {
		pointSet = append(pointSet, geo.Point{p.Lng, p.Lat})
	}
	return osrm.TableRequest{
		Profile:      "car",
		Coordinates:  osrm.NewGeometryFromPointSet(pointSet),
		Sources:      makeRange(0, len(in.Origins)-1),
		Destinations: makeRange(len(in.Origins), len(in.Origins)+len(in.Destinations)-1),
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func fromTableRes(res *osrm.TableResponse) *distance.MatrixResponse {
	var rows []distance.MatrixElementsRow
	for i := range res.Durations {
		origin := res.Durations[i]
		var elements []distance.MatrixElement
		for j := range origin {
			destination := origin[j]
			duration := time.Duration(destination) * time.Second
			elements = append(elements, distance.MatrixElement{
				Status:            "",
				Duration:          duration,
				DurationInTraffic: 0,
				Distance:          0,
			})
		}
		rows = append(rows, distance.MatrixElementsRow{Elements: elements})
	}
	return &distance.MatrixResponse{
		OriginAddresses:      nil,
		DestinationAddresses: nil,
		Rows:                 rows,
	}
}
