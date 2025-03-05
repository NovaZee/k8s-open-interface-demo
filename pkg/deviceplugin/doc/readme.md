## `DevicePlugin` 定义
### https://kubernetes.io/zh-cn/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/

这是 Device Plugin 向外暴露的 gRPC 服务接口，用于与 Kubernetes 的 Device Manager（设备管理器）交互。

### 方法总结

| 方法名                              | 输入参数                   | 返回值                        | 功能描述                                                                 |
|-------------------------------------|----------------------------|-------------------------------|-------------------------------------------------------------------------|
| `GetDevicePluginOptions`           | `Empty`                    | `DevicePluginOptions`         | 返回插件的选项，告诉 Device Manager 如何与插件通信。                     |
| `ListAndWatch`                     | `Empty`                    | `stream ListAndWatchResponse` | 以流的形式返回设备列表，设备状态变化或消失时更新并返回新列表。           |
| `GetPreferredAllocation`           | `PreferredAllocationRequest` | `PreferredAllocationResponse` | 从可用设备列表中返回推荐分配的设备，仅提供建议，最终由 Device Manager 决定。 |
| `Allocate`                         | `AllocateRequest`          | `AllocateResponse`            | 在容器创建时调用，插件执行设备特定操作并指示 kubelet 如何使设备可用。     |
| `PreStartContainer`                | `PreStartContainerRequest` | `PreStartContainerResponse`   | 若插件支持，在容器启动前调用，执行设备初始化操作（如重置设备）。         |

### 总体概述
- **核心功能**：定义了 Device Plugin 与 kubelet 的交互，包括设备状态报告（`ListAndWatch`）、设备分配（`Allocate`）及可选的初始化（`PreStartContainer`）。
- **通信方式**：基于 gRPC，使用流（`stream`）处理动态更新。
- **灵活性**：通过 `GetPreferredAllocation` 提供分配建议，但最终决策权在 Device Manager。