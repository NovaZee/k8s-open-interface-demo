package server

import (
	"context"
	pb "koid/uds/protoc/proto"
)

type CoreServer struct {
	pb.SyncDeviceServer
}

func (s *CoreServer) RunPodSandbox(ctx context.Context, req *pb.CheckDeviceStatus) (*pb.DeviceStatus, error) {
	return &pb.DeviceStatus{
		DeviceType:   "flex",
		DeviceName:   "prc1",
		DeviceStatus: "killed",
		DeviceMeta:   "{'local':''}",
	}, nil
}

func NewCoreServer() *CoreServer {
	return &CoreServer{}
}
