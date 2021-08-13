package system

import (
	"gin-starter/api/v1"
	"gin-starter/middleware"
	"github.com/gin-gonic/gin"
)

type JwtRouter struct {
}

func (s *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("jwt").Use(middleware.Operations())
	var jwtApi = v1.AppApi.SystemApi.Jwt
	{
		jwtRouter.POST("jsonInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
