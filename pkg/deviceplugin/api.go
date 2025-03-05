package deviceplugin

import (
	"fmt"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
	"log"
	"time"
)

func (p *IOPSDevicePlugin) report(s pluginapi.DevicePlugin_ListAndWatchServer) error {
	ticker := time.NewTicker(5 * time.Second) // 每 5 秒报告一次
	defer ticker.Stop()

	for {
		p.mu.Lock()
		resp := &pluginapi.ListAndWatchResponse{}
		for _, dev := range p.devices {
			resp.Devices = append(resp.Devices, &pluginapi.Device{
				ID:     dev.ID,
				Health: dev.Health,
			})
		}
		p.mu.Unlock()

		if err := s.Send(resp); err != nil {
			log.Printf("Failed to send device list: %v", err)
			return err
		}
		<-ticker.C
	}
}

func (p *IOPSDevicePlugin) allocate(req *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
	resp := &pluginapi.AllocateResponse{}
	p.mu.Lock()
	defer p.mu.Unlock()

	// 计算已分配的 IOPS 总数
	allocatedIOPS := 0
	for _, iops := range p.allocations {
		allocatedIOPS += iops
	}

	for _, containerReq := range req.ContainerRequests {
		for i, devID := range containerReq.DevicesIDs {
			if dev, ok := p.devices[devID]; ok && dev.Health == pluginapi.Healthy {
				requestedIOPS := 1000
				totalIOPS := p.currentIOPS + allocatedIOPS + requestedIOPS

				if totalIOPS > maxIOPS {
					log.Printf("Deny allocate %s: Now IOPS %d + Allocated %d + Request %d > Max %d",
						devID, p.currentIOPS, allocatedIOPS, requestedIOPS, maxIOPS)
					continue
				}

				// 记录分配
				podID := containerReq.DevicesIDs[i]
				if podID == "" {
					podID = fmt.Sprintf("pod-%s", devID) // 临时生成 Pod ID
				}
				p.allocations[podID] = requestedIOPS
				log.Printf("Allocate %s to Pod %s, IOPS: %d", devID, podID, requestedIOPS)

				resp.ContainerResponses = append(resp.ContainerResponses, &pluginapi.ContainerAllocateResponse{
					Envs: map[string]string{
						"STORAGE_IOPS_LIMIT": fmt.Sprintf("%d", requestedIOPS),
					},
					Mounts: []*pluginapi.Mount{
						{
							ContainerPath: "/sys/fs/cgroup/blkio",
							HostPath:      "/sys/fs/cgroup/blkio",
							ReadOnly:      false,
						},
					},
				})
			}
		}
	}
	return resp, nil
}
