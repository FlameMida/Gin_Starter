package system

import (
	"gin-starter/api/v1"
	"gin-starter/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.Operations())
	apiRouterWithoutRecord := Router.Group("api")
	var API = v1.AppApi.SystemApi.Api
	{
		apiRouter.POST("createApi", API.CreateApi)               // 创建Api
		apiRouter.POST("deleteApi", API.DeleteApi)               // 删除Api
		apiRouter.POST("getApiById", API.GetApiById)             // 获取单条Api消息
		apiRouter.POST("updateApi", API.UpdateApi)               // 更新api
		apiRouter.DELETE("deleteApisByIds", API.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", API.GetAllApis) // 获取所有api
		apiRouterWithoutRecord.POST("getApiList", API.GetApiList) // 获取Api列表
	}
}
