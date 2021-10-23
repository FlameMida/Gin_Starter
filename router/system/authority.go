package system

import (
	"gin-starter/api/v1"
	"gin-starter/middleware"
	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct {
}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.Operations())
	authorityRouterWithoutRecord := Router.Group("authority")
	var authorityApi = v1.AppApi.SystemApi.Authority
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority)   // 创建角色
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority)   // 删除角色
		authorityRouter.PUT("updateAuthority", authorityApi.UpdateAuthority)    // 更新角色
		authorityRouter.POST("copyAuthority", authorityApi.CopyAuthority)       // 更新角色
		authorityRouter.POST("setDataAuthority", authorityApi.SetDataAuthority) // 设置角色资源权限
	}
	{
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
