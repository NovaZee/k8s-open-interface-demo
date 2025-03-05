package server

import (
	"context"
	"koid/pkg/uds/protoc/proto"
)

type CoreServer struct {
	proto.SyncDeviceServer
}

func (s *CoreServer) GetDeviceStatus(context.Context, *proto.CheckDeviceStatus) (*proto.DeviceStatus, error) {
	return &proto.DeviceStatus{
		DeviceType:   "flex",
		DeviceName:   "prc1",
		DeviceStatus: "pending",
		DeviceMeta:   "{'local':''}",
	}, nil
}

func NewCoreServer() *CoreServer {
	return &CoreServer{}
}
