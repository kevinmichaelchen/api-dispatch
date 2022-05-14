package app

import (
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/kevinmichaelchen/api-dispatch/internal/distance"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"
	"googlemaps.github.io/maps"
	"os"
)

func NewMapsClient() (*maps.Client, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, errors.New("missing API_KEY for Google Maps")
	}
	c, err := maps.NewClient(
		maps.WithAPIKey(apiKey),
		maps.WithHTTPClient(otelhttp.DefaultClient),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build Google Maps client: %w", err)
	}
	return c, nil
}

func NewDistanceService(logger *zap.Logger, client *maps.Client) (*distance.Service, error) {
	if client == nil {
		return nil, errors.New("no maps client")
	}
	return distance.NewService(client), nil
}
