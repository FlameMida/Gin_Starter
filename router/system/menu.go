package system

import (
	"gin-starter/api/v1"
	"gin-starter/middleware"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.Operations())
	var authorityMenuApi = v1.AppApi.SystemApi.AuthorityMenu
	{
		menuRouter.POST("getMenu", authorityMenuApi.GetMenu)                   // 获取菜单树
		menuRouter.POST("getMenuList", authorityMenuApi.GetMenuList)           // 分页获取基础menu列表
		menuRouter.POST("addBaseMenu", authorityMenuApi.AddBaseMenu)           // 新增菜单
		menuRouter.POST("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)   // 获取用户动态路由
		menuRouter.POST("addMenuAuthority", authorityMenuApi.AddMenuAuthority) //	增加menu和角色关联关系
		menuRouter.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // 获取指定角色menu
		menuRouter.POST("deleteBaseMenu", authorityMenuApi.DeleteBaseMenu)     // 删除菜单
		menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)     // 更新菜单
		menuRouter.POST("getBaseMenuById", authorityMenuApi.GetBaseMenuById)   // 根据id获取菜单
	}
	return menuRouter
}
