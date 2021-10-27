package initialize

import (
	_ "gin-starter/docs"
	"gin-starter/global"
	"gin-starter/middleware"
	"gin-starter/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.CONFIG.Local.Path, http.Dir(global.CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	global.LOG.Info("use middleware logger")
	// 跨域
	// Router.Use(middleware.Cors())
	global.LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	// 获取路由组实例
	systemRouter := router.AppRouter.System
	exampleRouter := router.AppRouter.Example
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)             // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)             // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)            // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)            // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)          // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)          // 权限相关路由
		systemRouter.InitAuthorityRouter(PrivateGroup)       // 注册角色路由
		systemRouter.InitOperationsRouter(PrivateGroup)      // 操作记录
		exampleRouter.InitFileTransferRouter(PrivateGroup)   // 文件上传下载功能路由
		exampleRouter.InitExcelRouter(PrivateGroup)          // 表格导入导出
		exampleRouter.InitSimpleUploaderRouter(PrivateGroup) // 断点续传（插件版）
		exampleRouter.InitCustomerRouter(PrivateGroup)       // 客户路由

	}
	global.LOG.Info("router register success")
	return Router
}
