package system

import (
	"gin-starter/api/v1"
	"github.com/gin-gonic/gin"
)

type InitRouter struct {
}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	var dbApi = v1.AppApi.SystemApi.DB
	{
		initRouter.POST("initDB", dbApi.InitDB)   // 创建Api
		initRouter.POST("checkDB", dbApi.CheckDB) // 创建Api
	}
}
