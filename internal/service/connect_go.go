package service

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
)

// ConnectWrapper wraps our gRPC service.
type ConnectWrapper struct {
	s *Service
}

func NewConnectWrapper(s *Service) *ConnectWrapper {
	return &ConnectWrapper{s: s}
}

func (c *ConnectWrapper) UpdateDriverLocations(
	ctx context.Context,
	req *connect.Request[v1beta1.UpdateDriverLocationsRequest],
) (*connect.Response[v1beta1.UpdateDriverLocationsResponse], error) {
	res, err := c.s.UpdateDriverLocations(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")
	return out, nil
}

func (c *ConnectWrapper) CreateTrips(
	ctx context.Context,
	req *connect.Request[v1beta1.CreateTripsRequest],
) (*connect.Response[v1beta1.CreateTripsResponse], error) {
	res, err := c.s.CreateTrips(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")
	return out, nil
}

func (c *ConnectWrapper) GetNearestDrivers(
	ctx context.Context,
	req *connect.Request[v1beta1.GetNearestDriversRequest],
) (*connect.Response[v1beta1.GetNearestDriversResponse], error) {
	res, err := c.s.GetNearestDrivers(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")
	return out, nil
}

func (c *ConnectWrapper) GetNearestTrips(
	ctx context.Context,
	req *connect.Request[v1beta1.GetNearestTripsRequest],
) (*connect.Response[v1beta1.GetNearestTripsResponse], error) {
	res, err := c.s.GetNearestTrips(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")
	return out, nil
}

func (c *ConnectWrapper) ListDrivers(
	ctx context.Context,
	req *connect.Request[v1beta1.ListDriversRequest],
) (*connect.Response[v1beta1.ListDriversResponse], error) {
	res, err := c.s.ListDrivers(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")
	return out, nil
}
