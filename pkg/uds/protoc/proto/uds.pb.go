// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: uds.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckDeviceStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceType    int32                  `protobuf:"varint,1,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	DeviceName    string                 `protobuf:"bytes,2,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckDeviceStatus) Reset() {
	*x = CheckDeviceStatus{}
	mi := &file_uds_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckDeviceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckDeviceStatus) ProtoMessage() {}

func (x *CheckDeviceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_uds_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckDeviceStatus.ProtoReflect.Descriptor instead.
func (*CheckDeviceStatus) Descriptor() ([]byte, []int) {
	return file_uds_proto_rawDescGZIP(), []int{0}
}

func (x *CheckDeviceStatus) GetDeviceType() int32 {
	if x != nil {
		return x.DeviceType
	}
	return 0
}

func (x *CheckDeviceStatus) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

type DeviceStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceType    string                 `protobuf:"bytes,1,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	DeviceName    string                 `protobuf:"bytes,2,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	DeviceStatus  string                 `protobuf:"bytes,3,opt,name=deviceStatus,proto3" json:"deviceStatus,omitempty"`
	DeviceMeta    string                 `protobuf:"bytes,4,opt,name=deviceMeta,proto3" json:"deviceMeta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceStatus) Reset() {
	*x = DeviceStatus{}
	mi := &file_uds_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceStatus) ProtoMessage() {}

func (x *DeviceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_uds_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceStatus.ProtoReflect.Descriptor instead.
func (*DeviceStatus) Descriptor() ([]byte, []int) {
	return file_uds_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceStatus) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *DeviceStatus) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *DeviceStatus) GetDeviceStatus() string {
	if x != nil {
		return x.DeviceStatus
	}
	return ""
}

func (x *DeviceStatus) GetDeviceMeta() string {
	if x != nil {
		return x.DeviceMeta
	}
	return ""
}

var File_uds_proto protoreflect.FileDescriptor

var file_uds_proto_rawDesc = string([]byte{
	0x0a, 0x09, 0x75, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x32, 0x54,
	0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x15, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_uds_proto_rawDescOnce sync.Once
	file_uds_proto_rawDescData []byte
)

func file_uds_proto_rawDescGZIP() []byte {
	file_uds_proto_rawDescOnce.Do(func() {
		file_uds_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_uds_proto_rawDesc), len(file_uds_proto_rawDesc)))
	})
	return file_uds_proto_rawDescData
}

var file_uds_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_uds_proto_goTypes = []any{
	(*CheckDeviceStatus)(nil), // 0: greeter.CheckDeviceStatus
	(*DeviceStatus)(nil),      // 1: greeter.DeviceStatus
}
var file_uds_proto_depIdxs = []int32{
	0, // 0: greeter.SyncDevice.GetDeviceStatus:input_type -> greeter.CheckDeviceStatus
	1, // 1: greeter.SyncDevice.GetDeviceStatus:output_type -> greeter.DeviceStatus
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_uds_proto_init() }
func file_uds_proto_init() {
	if File_uds_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_uds_proto_rawDesc), len(file_uds_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_uds_proto_goTypes,
		DependencyIndexes: file_uds_proto_depIdxs,
		MessageInfos:      file_uds_proto_msgTypes,
	}.Build()
	File_uds_proto = out.File
	file_uds_proto_goTypes = nil
	file_uds_proto_depIdxs = nil
}
