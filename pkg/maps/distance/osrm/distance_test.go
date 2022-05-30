package osrm

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"net/http"
	"testing"
	"time"
)

func TestBetweenPoints(t *testing.T) {
	a := maps.LatLng{
		Lat: 40.791680675548136,
		Lng: -73.9650115649754,
	}
	b := maps.LatLng{
		Lat: 40.76866089218841,
		Lng: -73.98145413365043,
	}

	ctx := ctxzap.ToContext(context.Background(), zaptest.NewLogger(t))
	res, err := BetweenPoints(ctx, new(http.Client), distance.BetweenPointsInput{
		Destinations: []maps.LatLng{a},
		Origins:      []maps.LatLng{b},
	})
	require.NoError(t, err)
	require.Len(t, res.Rows, 1)
	require.Len(t, res.Rows[0].Elements, 1)
	require.Greater(t, res.Rows[0].Elements[0].Duration, time.Duration(0))
}
