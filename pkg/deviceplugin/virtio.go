package deviceplugin

import (
	"context"
	"google.golang.org/grpc"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

// IOPSDevicePlugin implements the Kubernetes device plugin API
type IOPSDevicePlugin struct {
	devices     map[string]*pluginapi.Device // 设备列表
	allocations map[string]int               // 记录分配的 IOPS，键为 Pod ID，值为分配的 IOPS
	currentIOPS int                          // 当前实时 IOPS
	server      *grpc.Server
	mu          sync.Mutex // 保护 devices 和 allocations
}

// NewIOPSDevicePlugin 初始化插件
func NewIOPSDevicePlugin() *IOPSDevicePlugin {
	devices := make(map[string]*pluginapi.Device)
	iopsMap := machineIO()
	for devID, _ := range iopsMap {
		devices[devID] = &pluginapi.Device{
			ID:     devID,
			Health: pluginapi.Healthy,
		}
	}

	plugin := &IOPSDevicePlugin{
		devices:     devices,
		currentIOPS: 0,
	}
	// 启动时从文件加载 allocations 到内存
	if err := plugin.loadAllocations(); err != nil {
		log.Printf("Load allocations failed: %v", err)
	}

	// sync IOPS
	go plugin.syncIOPS()

	return plugin
}

// machineIO 获取本机磁盘的 IOPS 信息
func machineIO() map[string]int {
	iopsMap := make(map[string]int)

	// use iostat fetch
	out, err := exec.Command("iostat", "-d", "1", "2").Output()
	if err == nil {
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			fields := strings.Fields(line)
			if len(fields) > 2 && strings.HasPrefix(fields[0], "vda") {
				rIOPS, _ := strconv.ParseFloat(fields[3], 64)
				wIOPS, _ := strconv.ParseFloat(fields[4], 64)
				totalIOPS := int(rIOPS + wIOPS)
				iopsMap[fields[0]] = totalIOPS
			}
		}
		return iopsMap
	}
	// fake
	iopsMap["vda"] = 10000
	return iopsMap
}

func (p *IOPSDevicePlugin) Run() error {
	return p.run()
}
func (p *IOPSDevicePlugin) Close() error {
	if err := p.saveAllocations(); err != nil {
		log.Printf("持久化 allocations 失败: %v", err)
	}
	if p.server != nil {
		p.server.Stop()
	}
	err := os.Remove(socketPath)
	if err != nil {
		return err
	}
	return nil
}

func (p *IOPSDevicePlugin) ListAndWatch(_ *pluginapi.Empty, s pluginapi.DevicePlugin_ListAndWatchServer) error {
	return p.report(s)
}

func (p *IOPSDevicePlugin) Allocate(_ context.Context, req *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
	return p.allocate(req)
}
func (p *IOPSDevicePlugin) GetDevicePluginOptions(context.Context, *pluginapi.Empty) (*pluginapi.DevicePluginOptions, error) {
	return &pluginapi.DevicePluginOptions{}, nil
}

func (p *IOPSDevicePlugin) PreStartContainer(context.Context, *pluginapi.PreStartContainerRequest) (*pluginapi.PreStartContainerResponse, error) {
	return &pluginapi.PreStartContainerResponse{}, nil
}

func (p *IOPSDevicePlugin) GetPreferredAllocation(context.Context, *pluginapi.PreferredAllocationRequest) (*pluginapi.PreferredAllocationResponse, error) {
	return &pluginapi.PreferredAllocationResponse{}, nil
}
