package osrm

import (
	"context"
	"errors"
	"fmt"
	osrm "github.com/gojuno/go.osrm"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	geo "github.com/paulmach/go.geo"
	"go.uber.org/zap"
	"time"
)

var (
	errFailedRequest = errors.New("failed OSRM request")
)

type CalculateOutput struct {
	Duration       time.Duration
	DistanceMeters int
}

func Calculate(ctx context.Context, a, b *v1beta1.LatLng) (*CalculateOutput, error) {
	logger := ctxzap.Extract(ctx)

	client := osrm.NewFromURL("https://router.project-osrm.org")

	res, err := client.Table(ctx, osrm.TableRequest{
		Profile: "car",
		Coordinates: osrm.NewGeometryFromPointSet(geo.PointSet{
			{a.GetLongitude(), a.GetLatitude()},
			{b.GetLongitude(), b.GetLatitude()},
		}),
		Sources:      []int{0},
		Destinations: []int{1},
	})

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

	return &CalculateOutput{
		Duration:       time.Duration(res.Durations[0][0]) * time.Second,
		DistanceMeters: 0,
	}, nil
}
