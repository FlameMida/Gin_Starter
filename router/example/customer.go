package example

import (
	v1 "gin-starter/api/v1"
	"gin-starter/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct {
}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.Operations())
	var customerApi = v1.AppApi.ExampleApi.CustomerApi
	{
		customerRouter.POST("customer", customerApi.CreateCustomer)     // 创建客户
		customerRouter.PUT("customer", customerApi.UpdateCustomer)      // 更新客户
		customerRouter.DELETE("customer", customerApi.DeleteCustomer)   // 删除客户
		customerRouter.GET("customer", customerApi.GetCustomer)         // 获取单一客户信息
		customerRouter.GET("customerList", customerApi.GetCustomerList) // 获取客户列表
	}
}
