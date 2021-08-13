package system

import (
	"gin-starter/api/v1"
	"gin-starter/middleware"
	"github.com/gin-gonic/gin"
)

type OperationsRouter struct {
}

func (s *OperationsRouter) InitOperationsRouter(Router *gin.RouterGroup) {
	operationsRouter := Router.Group("operations").Use(middleware.Operations())
	var operations = v1.AppApi.SystemApi.Operations
	{
		operationsRouter.POST("createOperations", operations.CreateOperations)             // 新建Operations
		operationsRouter.DELETE("deleteOperations", operations.DeleteOperations)           // 删除Operations
		operationsRouter.DELETE("deleteOperationsByIds", operations.DeleteOperationsByIds) // 批量删除Operations
		operationsRouter.GET("findOperations", operations.FindOperations)                  // 根据ID获取Operations
		operationsRouter.GET("getOperationsList", operations.GetOperationsList)            // 获取Operations列表

	}
}
