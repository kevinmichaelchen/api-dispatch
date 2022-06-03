// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: coop/drivers/dispatch/v1beta1/api.proto

package v1beta1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1beta1 "github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// DispatchServiceName is the fully-qualified name of the DispatchService service.
	DispatchServiceName = "coop.drivers.dispatch.v1beta1.DispatchService"
)

// DispatchServiceClient is a client for the coop.drivers.dispatch.v1beta1.DispatchService service.
type DispatchServiceClient interface {
	// Bulk-ingest driver locations.
	UpdateDriverLocations(context.Context, *connect_go.Request[v1beta1.UpdateDriverLocationsRequest]) (*connect_go.Response[v1beta1.UpdateDriverLocationsResponse], error)
	// Bulk-ingest (on-demand or scheduled) trips.
	CreateTrips(context.Context, *connect_go.Request[v1beta1.CreateTripsRequest]) (*connect_go.Response[v1beta1.CreateTripsResponse], error)
	// Gets the nearest drivers to a given trip pickup location.
	GetNearestDrivers(context.Context, *connect_go.Request[v1beta1.GetNearestDriversRequest]) (*connect_go.Response[v1beta1.GetNearestDriversResponse], error)
	// Gets the nearest trips to a given driver's location.
	GetNearestTrips(context.Context, *connect_go.Request[v1beta1.GetNearestTripsRequest]) (*connect_go.Response[v1beta1.GetNearestTripsResponse], error)
}

// NewDispatchServiceClient constructs a client for the
// coop.drivers.dispatch.v1beta1.DispatchService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDispatchServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) DispatchServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &dispatchServiceClient{
		updateDriverLocations: connect_go.NewClient[v1beta1.UpdateDriverLocationsRequest, v1beta1.UpdateDriverLocationsResponse](
			httpClient,
			baseURL+"/coop.drivers.dispatch.v1beta1.DispatchService/UpdateDriverLocations",
			opts...,
		),
		createTrips: connect_go.NewClient[v1beta1.CreateTripsRequest, v1beta1.CreateTripsResponse](
			httpClient,
			baseURL+"/coop.drivers.dispatch.v1beta1.DispatchService/CreateTrips",
			opts...,
		),
		getNearestDrivers: connect_go.NewClient[v1beta1.GetNearestDriversRequest, v1beta1.GetNearestDriversResponse](
			httpClient,
			baseURL+"/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestDrivers",
			opts...,
		),
		getNearestTrips: connect_go.NewClient[v1beta1.GetNearestTripsRequest, v1beta1.GetNearestTripsResponse](
			httpClient,
			baseURL+"/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestTrips",
			opts...,
		),
	}
}

// dispatchServiceClient implements DispatchServiceClient.
type dispatchServiceClient struct {
	updateDriverLocations *connect_go.Client[v1beta1.UpdateDriverLocationsRequest, v1beta1.UpdateDriverLocationsResponse]
	createTrips           *connect_go.Client[v1beta1.CreateTripsRequest, v1beta1.CreateTripsResponse]
	getNearestDrivers     *connect_go.Client[v1beta1.GetNearestDriversRequest, v1beta1.GetNearestDriversResponse]
	getNearestTrips       *connect_go.Client[v1beta1.GetNearestTripsRequest, v1beta1.GetNearestTripsResponse]
}

// UpdateDriverLocations calls coop.drivers.dispatch.v1beta1.DispatchService.UpdateDriverLocations.
func (c *dispatchServiceClient) UpdateDriverLocations(ctx context.Context, req *connect_go.Request[v1beta1.UpdateDriverLocationsRequest]) (*connect_go.Response[v1beta1.UpdateDriverLocationsResponse], error) {
	return c.updateDriverLocations.CallUnary(ctx, req)
}

// CreateTrips calls coop.drivers.dispatch.v1beta1.DispatchService.CreateTrips.
func (c *dispatchServiceClient) CreateTrips(ctx context.Context, req *connect_go.Request[v1beta1.CreateTripsRequest]) (*connect_go.Response[v1beta1.CreateTripsResponse], error) {
	return c.createTrips.CallUnary(ctx, req)
}

// GetNearestDrivers calls coop.drivers.dispatch.v1beta1.DispatchService.GetNearestDrivers.
func (c *dispatchServiceClient) GetNearestDrivers(ctx context.Context, req *connect_go.Request[v1beta1.GetNearestDriversRequest]) (*connect_go.Response[v1beta1.GetNearestDriversResponse], error) {
	return c.getNearestDrivers.CallUnary(ctx, req)
}

// GetNearestTrips calls coop.drivers.dispatch.v1beta1.DispatchService.GetNearestTrips.
func (c *dispatchServiceClient) GetNearestTrips(ctx context.Context, req *connect_go.Request[v1beta1.GetNearestTripsRequest]) (*connect_go.Response[v1beta1.GetNearestTripsResponse], error) {
	return c.getNearestTrips.CallUnary(ctx, req)
}

// DispatchServiceHandler is an implementation of the coop.drivers.dispatch.v1beta1.DispatchService
// service.
type DispatchServiceHandler interface {
	// Bulk-ingest driver locations.
	UpdateDriverLocations(context.Context, *connect_go.Request[v1beta1.UpdateDriverLocationsRequest]) (*connect_go.Response[v1beta1.UpdateDriverLocationsResponse], error)
	// Bulk-ingest (on-demand or scheduled) trips.
	CreateTrips(context.Context, *connect_go.Request[v1beta1.CreateTripsRequest]) (*connect_go.Response[v1beta1.CreateTripsResponse], error)
	// Gets the nearest drivers to a given trip pickup location.
	GetNearestDrivers(context.Context, *connect_go.Request[v1beta1.GetNearestDriversRequest]) (*connect_go.Response[v1beta1.GetNearestDriversResponse], error)
	// Gets the nearest trips to a given driver's location.
	GetNearestTrips(context.Context, *connect_go.Request[v1beta1.GetNearestTripsRequest]) (*connect_go.Response[v1beta1.GetNearestTripsResponse], error)
}

// NewDispatchServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDispatchServiceHandler(svc DispatchServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/coop.drivers.dispatch.v1beta1.DispatchService/UpdateDriverLocations", connect_go.NewUnaryHandler(
		"/coop.drivers.dispatch.v1beta1.DispatchService/UpdateDriverLocations",
		svc.UpdateDriverLocations,
		opts...,
	))
	mux.Handle("/coop.drivers.dispatch.v1beta1.DispatchService/CreateTrips", connect_go.NewUnaryHandler(
		"/coop.drivers.dispatch.v1beta1.DispatchService/CreateTrips",
		svc.CreateTrips,
		opts...,
	))
	mux.Handle("/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestDrivers", connect_go.NewUnaryHandler(
		"/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestDrivers",
		svc.GetNearestDrivers,
		opts...,
	))
	mux.Handle("/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestTrips", connect_go.NewUnaryHandler(
		"/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestTrips",
		svc.GetNearestTrips,
		opts...,
	))
	return "/coop.drivers.dispatch.v1beta1.DispatchService/", mux
}

// UnimplementedDispatchServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDispatchServiceHandler struct{}

func (UnimplementedDispatchServiceHandler) UpdateDriverLocations(context.Context, *connect_go.Request[v1beta1.UpdateDriverLocationsRequest]) (*connect_go.Response[v1beta1.UpdateDriverLocationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("coop.drivers.dispatch.v1beta1.DispatchService.UpdateDriverLocations is not implemented"))
}

func (UnimplementedDispatchServiceHandler) CreateTrips(context.Context, *connect_go.Request[v1beta1.CreateTripsRequest]) (*connect_go.Response[v1beta1.CreateTripsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("coop.drivers.dispatch.v1beta1.DispatchService.CreateTrips is not implemented"))
}

func (UnimplementedDispatchServiceHandler) GetNearestDrivers(context.Context, *connect_go.Request[v1beta1.GetNearestDriversRequest]) (*connect_go.Response[v1beta1.GetNearestDriversResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("coop.drivers.dispatch.v1beta1.DispatchService.GetNearestDrivers is not implemented"))
}

func (UnimplementedDispatchServiceHandler) GetNearestTrips(context.Context, *connect_go.Request[v1beta1.GetNearestTripsRequest]) (*connect_go.Response[v1beta1.GetNearestTripsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("coop.drivers.dispatch.v1beta1.DispatchService.GetNearestTrips is not implemented"))
}