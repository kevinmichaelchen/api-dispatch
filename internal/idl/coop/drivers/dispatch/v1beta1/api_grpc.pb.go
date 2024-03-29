// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: coop/drivers/dispatch/v1beta1/api.proto

package dispatchv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DispatchServiceClient is the client API for DispatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatchServiceClient interface {
	// Bulk-ingest driver locations.
	UpdateDriverLocations(ctx context.Context, in *UpdateDriverLocationsRequest, opts ...grpc.CallOption) (*UpdateDriverLocationsResponse, error)
	// Bulk-ingest (on-demand or scheduled) trips.
	CreateTrips(ctx context.Context, in *CreateTripsRequest, opts ...grpc.CallOption) (*CreateTripsResponse, error)
	// Gets the nearest drivers to a given trip pickup location.
	GetNearestDrivers(ctx context.Context, in *GetNearestDriversRequest, opts ...grpc.CallOption) (*GetNearestDriversResponse, error)
	// Gets the nearest trips to a given driver's location.
	GetNearestTrips(ctx context.Context, in *GetNearestTripsRequest, opts ...grpc.CallOption) (*GetNearestTripsResponse, error)
	// Lists drivers.
	ListDrivers(ctx context.Context, in *ListDriversRequest, opts ...grpc.CallOption) (*ListDriversResponse, error)
}

type dispatchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatchServiceClient(cc grpc.ClientConnInterface) DispatchServiceClient {
	return &dispatchServiceClient{cc}
}

func (c *dispatchServiceClient) UpdateDriverLocations(ctx context.Context, in *UpdateDriverLocationsRequest, opts ...grpc.CallOption) (*UpdateDriverLocationsResponse, error) {
	out := new(UpdateDriverLocationsResponse)
	err := c.cc.Invoke(ctx, "/coop.drivers.dispatch.v1beta1.DispatchService/UpdateDriverLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchServiceClient) CreateTrips(ctx context.Context, in *CreateTripsRequest, opts ...grpc.CallOption) (*CreateTripsResponse, error) {
	out := new(CreateTripsResponse)
	err := c.cc.Invoke(ctx, "/coop.drivers.dispatch.v1beta1.DispatchService/CreateTrips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchServiceClient) GetNearestDrivers(ctx context.Context, in *GetNearestDriversRequest, opts ...grpc.CallOption) (*GetNearestDriversResponse, error) {
	out := new(GetNearestDriversResponse)
	err := c.cc.Invoke(ctx, "/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestDrivers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchServiceClient) GetNearestTrips(ctx context.Context, in *GetNearestTripsRequest, opts ...grpc.CallOption) (*GetNearestTripsResponse, error) {
	out := new(GetNearestTripsResponse)
	err := c.cc.Invoke(ctx, "/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestTrips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatchServiceClient) ListDrivers(ctx context.Context, in *ListDriversRequest, opts ...grpc.CallOption) (*ListDriversResponse, error) {
	out := new(ListDriversResponse)
	err := c.cc.Invoke(ctx, "/coop.drivers.dispatch.v1beta1.DispatchService/ListDrivers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatchServiceServer is the server API for DispatchService service.
// All implementations should embed UnimplementedDispatchServiceServer
// for forward compatibility
type DispatchServiceServer interface {
	// Bulk-ingest driver locations.
	UpdateDriverLocations(context.Context, *UpdateDriverLocationsRequest) (*UpdateDriverLocationsResponse, error)
	// Bulk-ingest (on-demand or scheduled) trips.
	CreateTrips(context.Context, *CreateTripsRequest) (*CreateTripsResponse, error)
	// Gets the nearest drivers to a given trip pickup location.
	GetNearestDrivers(context.Context, *GetNearestDriversRequest) (*GetNearestDriversResponse, error)
	// Gets the nearest trips to a given driver's location.
	GetNearestTrips(context.Context, *GetNearestTripsRequest) (*GetNearestTripsResponse, error)
	// Lists drivers.
	ListDrivers(context.Context, *ListDriversRequest) (*ListDriversResponse, error)
}

// UnimplementedDispatchServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDispatchServiceServer struct {
}

func (UnimplementedDispatchServiceServer) UpdateDriverLocations(context.Context, *UpdateDriverLocationsRequest) (*UpdateDriverLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDriverLocations not implemented")
}
func (UnimplementedDispatchServiceServer) CreateTrips(context.Context, *CreateTripsRequest) (*CreateTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrips not implemented")
}
func (UnimplementedDispatchServiceServer) GetNearestDrivers(context.Context, *GetNearestDriversRequest) (*GetNearestDriversResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNearestDrivers not implemented")
}
func (UnimplementedDispatchServiceServer) GetNearestTrips(context.Context, *GetNearestTripsRequest) (*GetNearestTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNearestTrips not implemented")
}
func (UnimplementedDispatchServiceServer) ListDrivers(context.Context, *ListDriversRequest) (*ListDriversResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDrivers not implemented")
}

// UnsafeDispatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatchServiceServer will
// result in compilation errors.
type UnsafeDispatchServiceServer interface {
	mustEmbedUnimplementedDispatchServiceServer()
}

func RegisterDispatchServiceServer(s grpc.ServiceRegistrar, srv DispatchServiceServer) {
	s.RegisterService(&DispatchService_ServiceDesc, srv)
}

func _DispatchService_UpdateDriverLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDriverLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchServiceServer).UpdateDriverLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coop.drivers.dispatch.v1beta1.DispatchService/UpdateDriverLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchServiceServer).UpdateDriverLocations(ctx, req.(*UpdateDriverLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchService_CreateTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchServiceServer).CreateTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coop.drivers.dispatch.v1beta1.DispatchService/CreateTrips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchServiceServer).CreateTrips(ctx, req.(*CreateTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchService_GetNearestDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNearestDriversRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchServiceServer).GetNearestDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestDrivers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchServiceServer).GetNearestDrivers(ctx, req.(*GetNearestDriversRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchService_GetNearestTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNearestTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchServiceServer).GetNearestTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestTrips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchServiceServer).GetNearestTrips(ctx, req.(*GetNearestTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatchService_ListDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDriversRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchServiceServer).ListDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coop.drivers.dispatch.v1beta1.DispatchService/ListDrivers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchServiceServer).ListDrivers(ctx, req.(*ListDriversRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DispatchService_ServiceDesc is the grpc.ServiceDesc for DispatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DispatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coop.drivers.dispatch.v1beta1.DispatchService",
	HandlerType: (*DispatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateDriverLocations",
			Handler:    _DispatchService_UpdateDriverLocations_Handler,
		},
		{
			MethodName: "CreateTrips",
			Handler:    _DispatchService_CreateTrips_Handler,
		},
		{
			MethodName: "GetNearestDrivers",
			Handler:    _DispatchService_GetNearestDrivers_Handler,
		},
		{
			MethodName: "GetNearestTrips",
			Handler:    _DispatchService_GetNearestTrips_Handler,
		},
		{
			MethodName: "ListDrivers",
			Handler:    _DispatchService_ListDrivers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coop/drivers/dispatch/v1beta1/api.proto",
}
