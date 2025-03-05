package deviceplugin

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
	"log"
	"net"
	"os"
	"path"
	"time"
)

const (
	resourceName = "example.com/storage-iops"
	allocateFile = "/tmp/virtio.json"
	socketPath   = pluginapi.DevicePluginPath + "storage-iops.sock"
	maxIOPS      = 10000 // 定义最大 IOPS 值，可根据实际磁盘性能调整
)

func (p *IOPSDevicePlugin) run() error {
	os.Remove(socketPath)
	if err := os.MkdirAll(path.Dir(socketPath), 0755); err != nil {
		return fmt.Errorf("创建 socket 目录失败: %v", err)
	}

	lis, err := net.Listen("unix", socketPath)
	if err != nil {
		return fmt.Errorf("监听 socket 失败: %v", err)
	}

	p.server = grpc.NewServer()
	pluginapi.RegisterDevicePluginServer(p.server, p)

	go func() {
		if err := p.server.Serve(lis); err != nil {
			log.Fatalf("gRPC 服务启动失败: %v", err)
		}
	}()

	// 注册到 kubelet
	conn, err := grpc.Dial("unix:///var/lib/kubelet/device-plugins/kubelet.sock", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("连接 kubelet 失败: %v", err)
	}
	defer conn.Close()

	client := pluginapi.NewRegistrationClient(conn)
	_, err = client.Register(context.Background(), &pluginapi.RegisterRequest{
		Version:      pluginapi.Version,
		Endpoint:     "storage-iops.sock",
		ResourceName: resourceName,
	})
	if err != nil {
		return fmt.Errorf("注册到 kubelet 失败: %v", err)
	}

	log.Println("IOPS Device Plugin 启动成功")

	go p.saveAllocations()
	return nil
}

// syncIOPS 实时同步 vda 的 IOPS 到内存
func (p *IOPSDevicePlugin) syncIOPS() {
	ticker := time.NewTicker(1 * time.Second) // 每秒更新一次
	defer ticker.Stop()

	for range ticker.C {
		iopsMap := machineIO()
		p.mu.Lock()
		if iops, ok := iopsMap["vda"]; ok {
			p.currentIOPS = iops
			log.Printf("当前 vda IOPS: %d", p.currentIOPS)
		}
		// 更新设备状态
		for _, dev := range p.devices {
			if p.currentIOPS >= maxIOPS {
				dev.Health = pluginapi.Unhealthy
			} else {
				dev.Health = pluginapi.Healthy
			}
		}
		p.mu.Unlock()
	}
}

// saveAllocations 将 allocations 持久化到文件
func (p *IOPSDevicePlugin) saveAllocations() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := json.MarshalIndent(p.allocations, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化 allocations 失败: %v", err)
	}

	if err := os.WriteFile(allocateFile, data, 0644); err != nil {
		return fmt.Errorf("写入 allocations 文件失败: %v", err)
	}
	log.Printf("成功持久化 allocations 到 %s", allocateFile)
	return nil
}

// loadAllocations 从文件中加载 allocations 到内存
func (p *IOPSDevicePlugin) loadAllocations() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := os.ReadFile(allocateFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("文件 %s 不存在，初始化为空", allocateFile)
			return nil
		}
		return fmt.Errorf("读取文件 %s 失败: %v", allocateFile, err)
	}

	if err := json.Unmarshal(data, &p.allocations); err != nil {
		return fmt.Errorf("解析 allocations 文件失败: %v", err)
	}
	log.Printf("成功加载 allocations: %v", p.allocations)
	return nil
}
