package system

import (
	"gin-starter/api/v1"
	"gin-starter/middleware"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct {
}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Group("email").Use(middleware.Operations())
	var systemApi = v1.AppApi.SystemApi.Systems
	{
		emailRouter.POST("emailTest", systemApi.EmailTest) // 发送测试邮件
	}
}
