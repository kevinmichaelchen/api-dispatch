package osrm

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"net/http"
	"testing"
)

func TestReverseGeocode(t *testing.T) {
	g := NewGeocoder(new(http.Client))
	ctx := ctxzap.ToContext(context.Background(), zaptest.NewLogger(t))
	out, err := g.ReverseGeocode(ctx, maps.LatLng{Lat: -37.813611, Lng: 144.963056})
	require.NoError(t, err)
	pretty.Println(out)
}
