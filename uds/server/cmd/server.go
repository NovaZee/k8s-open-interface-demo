package main

import (
	"google.golang.org/grpc"
	pb "koid/uds/protoc/proto"
	"koid/uds/server"
	"log"
	"net"
	"os"
)

func main() {
	socketPath := "/tmp/uds.sock"

	if err := os.RemoveAll(socketPath); err != nil {
		log.Fatalf("Failed to remove socket: %v", err)
	}

	lis, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()
	pb.RegisterSyncDeviceServer(s, server.NewCoreServer())

	log.Printf("UDS Server listening on %s", socketPath)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
