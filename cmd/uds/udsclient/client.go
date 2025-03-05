package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"koid/pkg/uds/client"
	"log"
)

const protocol = "unix:///"

func main() {
	socketPath := "/tmp/uds.sock"

	conn, err := grpc.Dial(
		protocol+socketPath,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Control: Failed to connect to UDS: %v", err)
	}
	control := client.NewRemoteClient(conn)
	defer control.Close()
	resp, err := control.DeviceStatus(context.Background(), "prc1", 1)
	if err != nil {
		log.Fatalf("Control: Failed to call DeviceStatus: %v", err)
	}
	log.Printf("Control: DeviceStatus: %v", resp)
}
