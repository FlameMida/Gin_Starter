package system

import (
	"gin-starter/api/v1"
	"gin-starter/middleware"
	"github.com/gin-gonic/gin"
)

type SysRouter struct {
}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.Operations())
	var systems = v1.AppApi.SystemApi.Systems
	{
		sysRouter.POST("getSystemConfig", systems.GetSystemConfig) // 获取配置文件内容
		sysRouter.POST("setSystemConfig", systems.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("getServerInfo", systems.GetServerInfo)     // 获取服务器信息
		sysRouter.POST("reloadSystem", systems.ReloadSystem)       // 重启服务
	}
}
