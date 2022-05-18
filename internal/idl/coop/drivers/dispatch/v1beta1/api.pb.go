// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: coop/drivers/dispatch/v1beta1/api.proto

package v1beta1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateDriverLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*DriverLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *UpdateDriverLocationsRequest) Reset() {
	*x = UpdateDriverLocationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDriverLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDriverLocationsRequest) ProtoMessage() {}

func (x *UpdateDriverLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDriverLocationsRequest.ProtoReflect.Descriptor instead.
func (*UpdateDriverLocationsRequest) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateDriverLocationsRequest) GetLocations() []*DriverLocation {
	if x != nil {
		return x.Locations
	}
	return nil
}

type UpdateDriverLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDriverLocationsResponse) Reset() {
	*x = UpdateDriverLocationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDriverLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDriverLocationsResponse) ProtoMessage() {}

func (x *UpdateDriverLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDriverLocationsResponse.ProtoReflect.Descriptor instead.
func (*UpdateDriverLocationsResponse) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{1}
}

type CreateTripsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trips []*Trip `protobuf:"bytes,1,rep,name=trips,proto3" json:"trips,omitempty"`
}

func (x *CreateTripsRequest) Reset() {
	*x = CreateTripsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTripsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTripsRequest) ProtoMessage() {}

func (x *CreateTripsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTripsRequest.ProtoReflect.Descriptor instead.
func (*CreateTripsRequest) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTripsRequest) GetTrips() []*Trip {
	if x != nil {
		return x.Trips
	}
	return nil
}

type CreateTripsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTripsResponse) Reset() {
	*x = CreateTripsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTripsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTripsResponse) ProtoMessage() {}

func (x *CreateTripsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTripsResponse.ProtoReflect.Descriptor instead.
func (*CreateTripsResponse) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{3}
}

type GetNearestDriversRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// trip pickup location
	PickupLocation *LatLng `protobuf:"bytes,1,opt,name=pickup_location,json=pickupLocation,proto3" json:"pickup_location,omitempty"`
	Limit          int32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetNearestDriversRequest) Reset() {
	*x = GetNearestDriversRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNearestDriversRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNearestDriversRequest) ProtoMessage() {}

func (x *GetNearestDriversRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNearestDriversRequest.ProtoReflect.Descriptor instead.
func (*GetNearestDriversRequest) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetNearestDriversRequest) GetPickupLocation() *LatLng {
	if x != nil {
		return x.PickupLocation
	}
	return nil
}

func (x *GetNearestDriversRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetNearestDriversResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results       []*SearchResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	PickupAddress string          `protobuf:"bytes,2,opt,name=pickup_address,json=pickupAddress,proto3" json:"pickup_address,omitempty"`
}

func (x *GetNearestDriversResponse) Reset() {
	*x = GetNearestDriversResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNearestDriversResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNearestDriversResponse) ProtoMessage() {}

func (x *GetNearestDriversResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNearestDriversResponse.ProtoReflect.Descriptor instead.
func (*GetNearestDriversResponse) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetNearestDriversResponse) GetResults() []*SearchResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *GetNearestDriversResponse) GetPickupAddress() string {
	if x != nil {
		return x.PickupAddress
	}
	return ""
}

type GetNearestTripsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Driver:
	//	*GetNearestTripsRequest_DriverId
	//	*GetNearestTripsRequest_DriverLocation
	Driver isGetNearestTripsRequest_Driver `protobuf_oneof:"driver"`
	Limit  int32                           `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetNearestTripsRequest) Reset() {
	*x = GetNearestTripsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNearestTripsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNearestTripsRequest) ProtoMessage() {}

func (x *GetNearestTripsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNearestTripsRequest.ProtoReflect.Descriptor instead.
func (*GetNearestTripsRequest) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{6}
}

func (m *GetNearestTripsRequest) GetDriver() isGetNearestTripsRequest_Driver {
	if m != nil {
		return m.Driver
	}
	return nil
}

func (x *GetNearestTripsRequest) GetDriverId() string {
	if x, ok := x.GetDriver().(*GetNearestTripsRequest_DriverId); ok {
		return x.DriverId
	}
	return ""
}

func (x *GetNearestTripsRequest) GetDriverLocation() *LatLng {
	if x, ok := x.GetDriver().(*GetNearestTripsRequest_DriverLocation); ok {
		return x.DriverLocation
	}
	return nil
}

func (x *GetNearestTripsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type isGetNearestTripsRequest_Driver interface {
	isGetNearestTripsRequest_Driver()
}

type GetNearestTripsRequest_DriverId struct {
	DriverId string `protobuf:"bytes,1,opt,name=driver_id,json=driverId,proto3,oneof"`
}

type GetNearestTripsRequest_DriverLocation struct {
	DriverLocation *LatLng `protobuf:"bytes,2,opt,name=driver_location,json=driverLocation,proto3,oneof"`
}

func (*GetNearestTripsRequest_DriverId) isGetNearestTripsRequest_Driver() {}

func (*GetNearestTripsRequest_DriverLocation) isGetNearestTripsRequest_Driver() {}

type GetNearestTripsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SearchResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetNearestTripsResponse) Reset() {
	*x = GetNearestTripsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNearestTripsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNearestTripsResponse) ProtoMessage() {}

func (x *GetNearestTripsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNearestTripsResponse.ProtoReflect.Descriptor instead.
func (*GetNearestTripsResponse) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetNearestTripsResponse) GetResults() []*SearchResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type SearchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*SearchResult_Trip
	//	*SearchResult_Driver
	Payload isSearchResult_Payload `protobuf_oneof:"payload"`
	// Driver's distance from the pickup location (in meters)
	DistanceMeters float64 `protobuf:"fixed64,3,opt,name=distance_meters,json=distanceMeters,proto3" json:"distance_meters,omitempty"`
	// Time it takes driver to go to pickup
	Duration *durationpb.Duration `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	// The match's location.
	Location *LatLng `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	// Human-readable location
	Address string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	// The highest (finest) H3 resolution in which the match was found.
	Resolution int32 `protobuf:"varint,7,opt,name=resolution,proto3" json:"resolution,omitempty"`
	// The k-value of the lowest k-ring (most immediate neighborhood) in which the
	// match was found.
	KValue int32   `protobuf:"varint,8,opt,name=k_value,json=kValue,proto3" json:"k_value,omitempty"`
	Score  float64 `protobuf:"fixed64,9,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *SearchResult) Reset() {
	*x = SearchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResult) ProtoMessage() {}

func (x *SearchResult) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResult.ProtoReflect.Descriptor instead.
func (*SearchResult) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{8}
}

func (m *SearchResult) GetPayload() isSearchResult_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *SearchResult) GetTrip() *Trip {
	if x, ok := x.GetPayload().(*SearchResult_Trip); ok {
		return x.Trip
	}
	return nil
}

func (x *SearchResult) GetDriver() *DriverLocation {
	if x, ok := x.GetPayload().(*SearchResult_Driver); ok {
		return x.Driver
	}
	return nil
}

func (x *SearchResult) GetDistanceMeters() float64 {
	if x != nil {
		return x.DistanceMeters
	}
	return 0
}

func (x *SearchResult) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *SearchResult) GetLocation() *LatLng {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *SearchResult) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SearchResult) GetResolution() int32 {
	if x != nil {
		return x.Resolution
	}
	return 0
}

func (x *SearchResult) GetKValue() int32 {
	if x != nil {
		return x.KValue
	}
	return 0
}

func (x *SearchResult) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type isSearchResult_Payload interface {
	isSearchResult_Payload()
}

type SearchResult_Trip struct {
	Trip *Trip `protobuf:"bytes,1,opt,name=trip,proto3,oneof"`
}

type SearchResult_Driver struct {
	Driver *DriverLocation `protobuf:"bytes,2,opt,name=driver,proto3,oneof"`
}

func (*SearchResult_Trip) isSearchResult_Payload() {}

func (*SearchResult_Driver) isSearchResult_Payload() {}

var File_coop_drivers_dispatch_v1beta1_api_proto protoreflect.FileDescriptor

var file_coop_drivers_dispatch_v1beta1_api_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x6f, 0x70, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x28, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x74, 0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x1c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x09, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39,
	0x0a, 0x05, 0x74, 0x72, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x72,
	0x69, 0x70, 0x52, 0x05, 0x74, 0x72, 0x69, 0x70, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x80, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a,
	0x0f, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x0e, 0x70,
	0x69, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65,
	0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0xa9, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x54, 0x72,
	0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x0f, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xa9, 0x03,
	0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x39,
	0x0a, 0x04, 0x74, 0x72, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x72, 0x69,
	0x70, 0x48, 0x00, 0x52, 0x04, 0x74, 0x72, 0x69, 0x70, 0x12, 0x47, 0x0a, 0x06, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6f, 0x70,
	0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x64, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x0a, 0x07, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0xb0, 0x04, 0x0a, 0x0f, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x94, 0x01,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3b, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x69, 0x70, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69,
	0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x88, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f,
	0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x12, 0x35, 0x2e, 0x63, 0x6f,
	0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x54, 0x72, 0x69,
	0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x55, 0x5a, 0x53,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x76, 0x69, 0x6e,
	0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x63, 0x68, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coop_drivers_dispatch_v1beta1_api_proto_rawDescOnce sync.Once
	file_coop_drivers_dispatch_v1beta1_api_proto_rawDescData = file_coop_drivers_dispatch_v1beta1_api_proto_rawDesc
)

func file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP() []byte {
	file_coop_drivers_dispatch_v1beta1_api_proto_rawDescOnce.Do(func() {
		file_coop_drivers_dispatch_v1beta1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_coop_drivers_dispatch_v1beta1_api_proto_rawDescData)
	})
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescData
}

var file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_coop_drivers_dispatch_v1beta1_api_proto_goTypes = []interface{}{
	(*UpdateDriverLocationsRequest)(nil),  // 0: coop.drivers.dispatch.v1beta1.UpdateDriverLocationsRequest
	(*UpdateDriverLocationsResponse)(nil), // 1: coop.drivers.dispatch.v1beta1.UpdateDriverLocationsResponse
	(*CreateTripsRequest)(nil),            // 2: coop.drivers.dispatch.v1beta1.CreateTripsRequest
	(*CreateTripsResponse)(nil),           // 3: coop.drivers.dispatch.v1beta1.CreateTripsResponse
	(*GetNearestDriversRequest)(nil),      // 4: coop.drivers.dispatch.v1beta1.GetNearestDriversRequest
	(*GetNearestDriversResponse)(nil),     // 5: coop.drivers.dispatch.v1beta1.GetNearestDriversResponse
	(*GetNearestTripsRequest)(nil),        // 6: coop.drivers.dispatch.v1beta1.GetNearestTripsRequest
	(*GetNearestTripsResponse)(nil),       // 7: coop.drivers.dispatch.v1beta1.GetNearestTripsResponse
	(*SearchResult)(nil),                  // 8: coop.drivers.dispatch.v1beta1.SearchResult
	(*DriverLocation)(nil),                // 9: coop.drivers.dispatch.v1beta1.DriverLocation
	(*Trip)(nil),                          // 10: coop.drivers.dispatch.v1beta1.Trip
	(*LatLng)(nil),                        // 11: coop.drivers.dispatch.v1beta1.LatLng
	(*durationpb.Duration)(nil),           // 12: google.protobuf.Duration
}
var file_coop_drivers_dispatch_v1beta1_api_proto_depIdxs = []int32{
	9,  // 0: coop.drivers.dispatch.v1beta1.UpdateDriverLocationsRequest.locations:type_name -> coop.drivers.dispatch.v1beta1.DriverLocation
	10, // 1: coop.drivers.dispatch.v1beta1.CreateTripsRequest.trips:type_name -> coop.drivers.dispatch.v1beta1.Trip
	11, // 2: coop.drivers.dispatch.v1beta1.GetNearestDriversRequest.pickup_location:type_name -> coop.drivers.dispatch.v1beta1.LatLng
	8,  // 3: coop.drivers.dispatch.v1beta1.GetNearestDriversResponse.results:type_name -> coop.drivers.dispatch.v1beta1.SearchResult
	11, // 4: coop.drivers.dispatch.v1beta1.GetNearestTripsRequest.driver_location:type_name -> coop.drivers.dispatch.v1beta1.LatLng
	8,  // 5: coop.drivers.dispatch.v1beta1.GetNearestTripsResponse.results:type_name -> coop.drivers.dispatch.v1beta1.SearchResult
	10, // 6: coop.drivers.dispatch.v1beta1.SearchResult.trip:type_name -> coop.drivers.dispatch.v1beta1.Trip
	9,  // 7: coop.drivers.dispatch.v1beta1.SearchResult.driver:type_name -> coop.drivers.dispatch.v1beta1.DriverLocation
	12, // 8: coop.drivers.dispatch.v1beta1.SearchResult.duration:type_name -> google.protobuf.Duration
	11, // 9: coop.drivers.dispatch.v1beta1.SearchResult.location:type_name -> coop.drivers.dispatch.v1beta1.LatLng
	0,  // 10: coop.drivers.dispatch.v1beta1.DispatchService.UpdateDriverLocations:input_type -> coop.drivers.dispatch.v1beta1.UpdateDriverLocationsRequest
	2,  // 11: coop.drivers.dispatch.v1beta1.DispatchService.CreateTrips:input_type -> coop.drivers.dispatch.v1beta1.CreateTripsRequest
	4,  // 12: coop.drivers.dispatch.v1beta1.DispatchService.GetNearestDrivers:input_type -> coop.drivers.dispatch.v1beta1.GetNearestDriversRequest
	6,  // 13: coop.drivers.dispatch.v1beta1.DispatchService.GetNearestTrips:input_type -> coop.drivers.dispatch.v1beta1.GetNearestTripsRequest
	1,  // 14: coop.drivers.dispatch.v1beta1.DispatchService.UpdateDriverLocations:output_type -> coop.drivers.dispatch.v1beta1.UpdateDriverLocationsResponse
	3,  // 15: coop.drivers.dispatch.v1beta1.DispatchService.CreateTrips:output_type -> coop.drivers.dispatch.v1beta1.CreateTripsResponse
	5,  // 16: coop.drivers.dispatch.v1beta1.DispatchService.GetNearestDrivers:output_type -> coop.drivers.dispatch.v1beta1.GetNearestDriversResponse
	7,  // 17: coop.drivers.dispatch.v1beta1.DispatchService.GetNearestTrips:output_type -> coop.drivers.dispatch.v1beta1.GetNearestTripsResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_coop_drivers_dispatch_v1beta1_api_proto_init() }
func file_coop_drivers_dispatch_v1beta1_api_proto_init() {
	if File_coop_drivers_dispatch_v1beta1_api_proto != nil {
		return
	}
	file_coop_drivers_dispatch_v1beta1_driver_proto_init()
	file_coop_drivers_dispatch_v1beta1_latlng_proto_init()
	file_coop_drivers_dispatch_v1beta1_trip_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDriverLocationsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDriverLocationsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTripsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTripsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNearestDriversRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNearestDriversResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNearestTripsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNearestTripsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*GetNearestTripsRequest_DriverId)(nil),
		(*GetNearestTripsRequest_DriverLocation)(nil),
	}
	file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*SearchResult_Trip)(nil),
		(*SearchResult_Driver)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coop_drivers_dispatch_v1beta1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coop_drivers_dispatch_v1beta1_api_proto_goTypes,
		DependencyIndexes: file_coop_drivers_dispatch_v1beta1_api_proto_depIdxs,
		MessageInfos:      file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes,
	}.Build()
	File_coop_drivers_dispatch_v1beta1_api_proto = out.File
	file_coop_drivers_dispatch_v1beta1_api_proto_rawDesc = nil
	file_coop_drivers_dispatch_v1beta1_api_proto_goTypes = nil
	file_coop_drivers_dispatch_v1beta1_api_proto_depIdxs = nil
}
