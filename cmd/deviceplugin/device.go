package main

import (
	"koid/pkg/deviceplugin"
	"log"
)

func main() {
	plugin := deviceplugin.NewIOPSDevicePlugin()
	if err := plugin.Run(); err != nil {
		log.Fatalf("插件启动失败: %v", err)
	}
	defer plugin.Close()

	select {}
}
