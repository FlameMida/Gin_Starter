package system

import (
	"gin-starter/api/v1"
	"gin-starter/middleware"
	"github.com/gin-gonic/gin"
)

type CasbinRouter struct {
}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.Operations())
	casbinRouterWithoutRecord := Router.Group("casbin")
	var casbinApi = v1.AppApi.SystemApi.Casbin
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)

	}
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
