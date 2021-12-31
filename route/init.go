package route

import (
	"Learn-CasaOS/model"
	"Learn-CasaOS/model/system_app"
	"Learn-CasaOS/pkg/config"
	"Learn-CasaOS/pkg/docker"
	"Learn-CasaOS/pkg/utils/file"
	"Learn-CasaOS/service"
	"encoding/json"
	"time"
)

// check if the system application is installed
func checkSystemApp() {
	list := service.MyService.App().GetSystemAppList()
	for _, v := range *list {
		if v.Image == "linuxserver/syncthing" {
			if v.State != "running" {
				service.MyService.Docker().DockerContainerStart(v.CustomId)
			}
			syncIsExistence := true
			if config.SystemConfigInfo.SyncPort != v.Port {
				config.SystemConfigInfo.SyncPort = v.Port
			}
			var paths []model.PathMap
			json.Unmarshal([]byte(v.Volumes), &paths)
			path := ""
			for _, i := range paths {
				if i.ContainerPath == "/config" {
					path = docker.GetDir(v.CustomId, i.Path) + "config.xml"
					for i := 0; i < 10; i++ {
						if file.CheckNotExist(path) {
							time.Sleep(1 * time.Second)
						} else {
							break
						}
					}
					break
				}
			}
			content := file.ReadFullFile(path)
			syncConfig := &system_app.SyncConfig{}
		}
	}
}
