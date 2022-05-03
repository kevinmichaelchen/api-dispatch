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
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DriverLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId  string                 `protobuf:"bytes,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LatLng    *LatLng                `protobuf:"bytes,3,opt,name=lat_lng,json=latLng,proto3" json:"lat_lng,omitempty"`
}

func (x *DriverLocation) Reset() {
	*x = DriverLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverLocation) ProtoMessage() {}

func (x *DriverLocation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DriverLocation.ProtoReflect.Descriptor instead.
func (*DriverLocation) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{0}
}

func (x *DriverLocation) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *DriverLocation) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *DriverLocation) GetLatLng() *LatLng {
	if x != nil {
		return x.LatLng
	}
	return nil
}

type IngestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*DriverLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *IngestRequest) Reset() {
	*x = IngestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestRequest) ProtoMessage() {}

func (x *IngestRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IngestRequest.ProtoReflect.Descriptor instead.
func (*IngestRequest) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{1}
}

func (x *IngestRequest) GetLocations() []*DriverLocation {
	if x != nil {
		return x.Locations
	}
	return nil
}

type IngestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IngestResponse) Reset() {
	*x = IngestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestResponse) ProtoMessage() {}

func (x *IngestResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IngestResponse.ProtoReflect.Descriptor instead.
func (*IngestResponse) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{2}
}

type DispatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *LatLng `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Limit    int32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *DispatchRequest) Reset() {
	*x = DispatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchRequest) ProtoMessage() {}

func (x *DispatchRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DispatchRequest.ProtoReflect.Descriptor instead.
func (*DispatchRequest) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{3}
}

func (x *DispatchRequest) GetLocation() *LatLng {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *DispatchRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type DispatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results       []*SearchResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	PickupAddress string          `protobuf:"bytes,2,opt,name=pickup_address,json=pickupAddress,proto3" json:"pickup_address,omitempty"`
}

func (x *DispatchResponse) Reset() {
	*x = DispatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchResponse) ProtoMessage() {}

func (x *DispatchResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DispatchResponse.ProtoReflect.Descriptor instead.
func (*DispatchResponse) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{4}
}

func (x *DispatchResponse) GetResults() []*SearchResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *DispatchResponse) GetPickupAddress() string {
	if x != nil {
		return x.PickupAddress
	}
	return ""
}

type SearchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Driver's ID
	DriverId string `protobuf:"bytes,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	// Driver's distance from the pickup location (in meters)
	DistanceMeters float64 `protobuf:"fixed64,2,opt,name=distance_meters,json=distanceMeters,proto3" json:"distance_meters,omitempty"`
	// Time it takes driver to go to pickup
	Duration *durationpb.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	// Driver's current location
	DriverLocation *LatLng `protobuf:"bytes,4,opt,name=driver_location,json=driverLocation,proto3" json:"driver_location,omitempty"`
	// Driver's current location as a human-readable address
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// what resolution the driver was found to be a k=1 k-ring neighbor to the
	// pickup location. if resolution is 7, that means the highest (finest)
	// resolution in which the driver appeared in the pickup location's
	// neighborhood was 7.
	Resolution int32 `protobuf:"varint,6,opt,name=resolution,proto3" json:"resolution,omitempty"`
	// if k=1, it's an immediate (1st degree) neighbor. if k=2, it's a 2nd-degree
	// neighbor (i.e., a neighbor's neighbor).
	KValue int32 `protobuf:"varint,7,opt,name=k_value,json=kValue,proto3" json:"k_value,omitempty"`
}

func (x *SearchResult) Reset() {
	*x = SearchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResult) ProtoMessage() {}

func (x *SearchResult) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SearchResult.ProtoReflect.Descriptor instead.
func (*SearchResult) Descriptor() ([]byte, []int) {
	return file_coop_drivers_dispatch_v1beta1_api_proto_rawDescGZIP(), []int{5}
}

func (x *SearchResult) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
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

func (x *SearchResult) GetDriverLocation() *LatLng {
	if x != nil {
		return x.DriverLocation
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

var File_coop_drivers_dispatch_v1beta1_api_proto protoreflect.FileDescriptor

var file_coop_drivers_dispatch_v1beta1_api_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x6f, 0x70, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x3e, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x5f, 0x6c, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e,
	0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x22,
	0x5c, 0x0a, 0x0d, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4b, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x10, 0x0a,
	0x0e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x6a, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x10,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x45, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xae,
	0x02, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x0f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x0e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0xe9, 0x01, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x06, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x2e,
	0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6f,
	0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x08,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x55, 0x5a, 0x53, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x76, 0x69, 0x6e, 0x6d,
	0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x63, 0x68, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_coop_drivers_dispatch_v1beta1_api_proto_goTypes = []interface{}{
	(*DriverLocation)(nil),        // 0: coop.drivers.dispatch.v1beta1.DriverLocation
	(*IngestRequest)(nil),         // 1: coop.drivers.dispatch.v1beta1.IngestRequest
	(*IngestResponse)(nil),        // 2: coop.drivers.dispatch.v1beta1.IngestResponse
	(*DispatchRequest)(nil),       // 3: coop.drivers.dispatch.v1beta1.DispatchRequest
	(*DispatchResponse)(nil),      // 4: coop.drivers.dispatch.v1beta1.DispatchResponse
	(*SearchResult)(nil),          // 5: coop.drivers.dispatch.v1beta1.SearchResult
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*LatLng)(nil),                // 7: coop.drivers.dispatch.v1beta1.LatLng
	(*durationpb.Duration)(nil),   // 8: google.protobuf.Duration
}
var file_coop_drivers_dispatch_v1beta1_api_proto_depIdxs = []int32{
	6, // 0: coop.drivers.dispatch.v1beta1.DriverLocation.timestamp:type_name -> google.protobuf.Timestamp
	7, // 1: coop.drivers.dispatch.v1beta1.DriverLocation.lat_lng:type_name -> coop.drivers.dispatch.v1beta1.LatLng
	0, // 2: coop.drivers.dispatch.v1beta1.IngestRequest.locations:type_name -> coop.drivers.dispatch.v1beta1.DriverLocation
	7, // 3: coop.drivers.dispatch.v1beta1.DispatchRequest.location:type_name -> coop.drivers.dispatch.v1beta1.LatLng
	5, // 4: coop.drivers.dispatch.v1beta1.DispatchResponse.results:type_name -> coop.drivers.dispatch.v1beta1.SearchResult
	8, // 5: coop.drivers.dispatch.v1beta1.SearchResult.duration:type_name -> google.protobuf.Duration
	7, // 6: coop.drivers.dispatch.v1beta1.SearchResult.driver_location:type_name -> coop.drivers.dispatch.v1beta1.LatLng
	1, // 7: coop.drivers.dispatch.v1beta1.DispatchService.Ingest:input_type -> coop.drivers.dispatch.v1beta1.IngestRequest
	3, // 8: coop.drivers.dispatch.v1beta1.DispatchService.Dispatch:input_type -> coop.drivers.dispatch.v1beta1.DispatchRequest
	2, // 9: coop.drivers.dispatch.v1beta1.DispatchService.Ingest:output_type -> coop.drivers.dispatch.v1beta1.IngestResponse
	4, // 10: coop.drivers.dispatch.v1beta1.DispatchService.Dispatch:output_type -> coop.drivers.dispatch.v1beta1.DispatchResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_coop_drivers_dispatch_v1beta1_api_proto_init() }
func file_coop_drivers_dispatch_v1beta1_api_proto_init() {
	if File_coop_drivers_dispatch_v1beta1_api_proto != nil {
		return
	}
	file_coop_drivers_dispatch_v1beta1_latlng_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coop_drivers_dispatch_v1beta1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverLocation); i {
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
			switch v := v.(*IngestRequest); i {
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
			switch v := v.(*IngestResponse); i {
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
			switch v := v.(*DispatchRequest); i {
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
			switch v := v.(*DispatchResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coop_drivers_dispatch_v1beta1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
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
