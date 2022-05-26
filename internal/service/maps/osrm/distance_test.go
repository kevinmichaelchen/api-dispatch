package osrm

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"testing"
	"time"
)

func TestCalculate(t *testing.T) {
	a := &v1beta1.LatLng{
		Latitude:  40.791680675548136,
		Longitude: -73.9650115649754,
	}
	b := &v1beta1.LatLng{
		Latitude:  40.76866089218841,
		Longitude: -73.98145413365043,
	}

	ctx := ctxzap.ToContext(context.Background(), zaptest.NewLogger(t))
	res, err := Calculate(ctx, a, b)
	require.NoError(t, err)
	require.Greater(t, res.Duration, time.Duration(0))
}
