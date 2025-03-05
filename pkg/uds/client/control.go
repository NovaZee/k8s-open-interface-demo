package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"koid/pkg/uds/protoc/proto"
	"sync"
	"time"
)

var localCache = make(map[string]proto.DeviceStatus)
var mu sync.Mutex

type PluginCore struct {
	remoteClient proto.SyncDeviceClient
	timeout      time.Duration
	ClientConn   *grpc.ClientConn
}

func NewRemoteClient(conn *grpc.ClientConn) *PluginCore {
	c := proto.NewSyncDeviceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	localCache = make(map[string]proto.DeviceStatus)
	return &PluginCore{
		remoteClient: c,
		timeout:      5 * time.Second,
		ClientConn:   conn,
	}
}

type Device interface {
	DeviceStatus(ctx context.Context, deviceName string, kind int) (string, error)
}

func (p *PluginCore) DeviceStatus(ctx context.Context, deviceName string, kind int) (string, error) {
	connCtx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	status, f, err := p.deviceStatus(connCtx, deviceName, int(kind))
	if err != nil {
		return "", err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		loopCtx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
		defer wg.Done()
		defer cancel()
		_, err := f(loopCtx, deviceName)
		if err != nil {
			fmt.Println("Loop check failed: ", err)
		}
	}()
	wg.Wait()
	return status, nil
}

func (p *PluginCore) deviceStatus(ctx context.Context, deviceName string, kind int) (string, func(ctx context.Context, addr string) (string, error), error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	mu.Lock()
	defer mu.Unlock()

	status, err := p.remoteClient.GetDeviceStatus(ctx, &proto.CheckDeviceStatus{
		DeviceName: deviceName,
		DeviceType: int32(kind),
	})
	if err != nil {
		return "", nil, err
	}

	localCache[deviceName] = *status
	return status.GetDeviceStatus(), loopCheck, nil
}

func (p *PluginCore) Close() {
	p.ClientConn.Close()
}

// loopCheck
func loopCheck(ctx context.Context, deviceName string) (string, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-ticker.C:
			s, ok := localCache[deviceName]
			if !ok {
				time.Sleep(5 * time.Second)
				continue
			}
			fmt.Println(fmt.Sprintf("Device %s status: %s", deviceName, s.DeviceStatus))
		}
	}
}
