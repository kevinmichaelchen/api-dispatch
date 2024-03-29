// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: coop/drivers/dispatch/v1beta1/driver.proto

package dispatchv1beta1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

	Id                  string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DriverId            string                 `protobuf:"bytes,2,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	MostRecentHeartbeat *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=most_recent_heartbeat,json=mostRecentHeartbeat,proto3" json:"most_recent_heartbeat,omitempty"`
	CurrentLocation     *LatLng                `protobuf:"bytes,4,opt,name=current_location,json=currentLocation,proto3" json:"current_location,omitempty"`
}

func (x *DriverLocation) Reset() {
	*x = DriverLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_dispatch_v1beta1_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverLocation) ProtoMessage() {}

func (x *DriverLocation) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_dispatch_v1beta1_driver_proto_msgTypes[0]
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
	return file_coop_drivers_dispatch_v1beta1_driver_proto_rawDescGZIP(), []int{0}
}

func (x *DriverLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DriverLocation) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *DriverLocation) GetMostRecentHeartbeat() *timestamppb.Timestamp {
	if x != nil {
		return x.MostRecentHeartbeat
	}
	return nil
}

func (x *DriverLocation) GetCurrentLocation() *LatLng {
	if x != nil {
		return x.CurrentLocation
	}
	return nil
}

var File_coop_drivers_dispatch_v1beta1_driver_proto protoreflect.FileDescriptor

var file_coop_drivers_dispatch_v1beta1_driver_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f,
	0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6f,
	0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x0e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x58, 0x0a, 0x15, 0x6d, 0x6f,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x13, 0x6d, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x12, 0x5a, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c,
	0x61, 0x74, 0x4c, 0x6e, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0xac, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0b, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x65, 0x76, 0x69, 0x6e, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x63, 0x68,
	0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x6f,
	0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x44, 0x44,
	0xaa, 0x02, 0x1d, 0x43, 0x6f, 0x6f, 0x70, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x1d, 0x43, 0x6f, 0x6f, 0x70, 0x5c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x5c,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xe2, 0x02, 0x29, 0x43, 0x6f, 0x6f, 0x70, 0x5c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x5c,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x43,
	0x6f, 0x6f, 0x70, 0x3a, 0x3a, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x3a, 0x3a, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coop_drivers_dispatch_v1beta1_driver_proto_rawDescOnce sync.Once
	file_coop_drivers_dispatch_v1beta1_driver_proto_rawDescData = file_coop_drivers_dispatch_v1beta1_driver_proto_rawDesc
)

func file_coop_drivers_dispatch_v1beta1_driver_proto_rawDescGZIP() []byte {
	file_coop_drivers_dispatch_v1beta1_driver_proto_rawDescOnce.Do(func() {
		file_coop_drivers_dispatch_v1beta1_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_coop_drivers_dispatch_v1beta1_driver_proto_rawDescData)
	})
	return file_coop_drivers_dispatch_v1beta1_driver_proto_rawDescData
}

var file_coop_drivers_dispatch_v1beta1_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_coop_drivers_dispatch_v1beta1_driver_proto_goTypes = []interface{}{
	(*DriverLocation)(nil),        // 0: coop.drivers.dispatch.v1beta1.DriverLocation
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*LatLng)(nil),                // 2: coop.drivers.dispatch.v1beta1.LatLng
}
var file_coop_drivers_dispatch_v1beta1_driver_proto_depIdxs = []int32{
	1, // 0: coop.drivers.dispatch.v1beta1.DriverLocation.most_recent_heartbeat:type_name -> google.protobuf.Timestamp
	2, // 1: coop.drivers.dispatch.v1beta1.DriverLocation.current_location:type_name -> coop.drivers.dispatch.v1beta1.LatLng
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coop_drivers_dispatch_v1beta1_driver_proto_init() }
func file_coop_drivers_dispatch_v1beta1_driver_proto_init() {
	if File_coop_drivers_dispatch_v1beta1_driver_proto != nil {
		return
	}
	file_coop_drivers_dispatch_v1beta1_latlng_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coop_drivers_dispatch_v1beta1_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coop_drivers_dispatch_v1beta1_driver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coop_drivers_dispatch_v1beta1_driver_proto_goTypes,
		DependencyIndexes: file_coop_drivers_dispatch_v1beta1_driver_proto_depIdxs,
		MessageInfos:      file_coop_drivers_dispatch_v1beta1_driver_proto_msgTypes,
	}.Build()
	File_coop_drivers_dispatch_v1beta1_driver_proto = out.File
	file_coop_drivers_dispatch_v1beta1_driver_proto_rawDesc = nil
	file_coop_drivers_dispatch_v1beta1_driver_proto_goTypes = nil
	file_coop_drivers_dispatch_v1beta1_driver_proto_depIdxs = nil
}
