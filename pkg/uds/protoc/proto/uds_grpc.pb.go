// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: uds.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SyncDevice_GetDeviceStatus_FullMethodName = "/greeter.SyncDevice/GetDeviceStatus"
)

// SyncDeviceClient is the server API for SyncDevice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncDeviceClient interface {
	GetDeviceStatus(ctx context.Context, in *CheckDeviceStatus, opts ...grpc.CallOption) (*DeviceStatus, error)
}

type syncDeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncDeviceClient(cc grpc.ClientConnInterface) SyncDeviceClient {
	return &syncDeviceClient{cc}
}

func (c *syncDeviceClient) GetDeviceStatus(ctx context.Context, in *CheckDeviceStatus, opts ...grpc.CallOption) (*DeviceStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeviceStatus)
	err := c.cc.Invoke(ctx, SyncDevice_GetDeviceStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncDeviceServer is the client API for SyncDevice service.
// All implementations must embed UnimplementedSyncDeviceServer
// for forward compatibility.
type SyncDeviceServer interface {
	GetDeviceStatus(context.Context, *CheckDeviceStatus) (*DeviceStatus, error)
	mustEmbedUnimplementedSyncDeviceServer()
}

// UnimplementedSyncDeviceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSyncDeviceServer struct{}

func (UnimplementedSyncDeviceServer) GetDeviceStatus(context.Context, *CheckDeviceStatus) (*DeviceStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceStatus not implemented")
}
func (UnimplementedSyncDeviceServer) mustEmbedUnimplementedSyncDeviceServer() {}
func (UnimplementedSyncDeviceServer) testEmbeddedByValue()                    {}

// UnsafeSyncDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncDeviceServer will
// result in compilation errors.
type UnsafeSyncDeviceServer interface {
	mustEmbedUnimplementedSyncDeviceServer()
}

func RegisterSyncDeviceServer(s grpc.ServiceRegistrar, srv SyncDeviceServer) {
	// If the following call pancis, it indicates UnimplementedSyncDeviceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SyncDevice_ServiceDesc, srv)
}

func _SyncDevice_GetDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDeviceStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncDeviceServer).GetDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncDevice_GetDeviceStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncDeviceServer).GetDeviceStatus(ctx, req.(*CheckDeviceStatus))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncDevice_ServiceDesc is the grpc.ServiceDesc for SyncDevice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncDevice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.SyncDevice",
	HandlerType: (*SyncDeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceStatus",
			Handler:    _SyncDevice_GetDeviceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uds.proto",
}
