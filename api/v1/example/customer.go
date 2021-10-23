package example

import (
	"gin-starter/global"
	"gin-starter/model/common/request"
	"gin-starter/model/common/response"
	"gin-starter/model/example"
	exampleRes "gin-starter/model/example/response"
	"gin-starter/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerApi struct {
}

// CreateCustomer
//
// @Tags Customer
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.Customer true "客户用户名, 客户手机号码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /customer/customer [post]
func (e *CustomerApi) CreateCustomer(c *gin.Context) {
	var customer example.Customer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customer.UserID = utils.GetUserID(c)
	customer.UserAuthorityID = utils.GetUserAuthorityId(c)
	if err := customerService.CreateCustomer(customer); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCustomer
// @Tags Customer
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.Customer true "客户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customer/customer [delete]
func (e *CustomerApi) DeleteCustomer(c *gin.Context) {
	var customer example.Customer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := customerService.DeleteCustomer(customer); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateCustomer
// @Tags Customer
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.Customer true "客户ID, 客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customer/customer [put]
func (e *CustomerApi) UpdateCustomer(c *gin.Context) {
	var customer example.Customer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := customerService.UpdateCustomer(&customer); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetCustomer
// @Tags Customer
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body example.Customer true "客户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
func (e *CustomerApi) GetCustomer(c *gin.Context) {
	var customer example.Customer
	_ = c.ShouldBindQuery(&customer)
	if err := utils.Verify(customer.MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, data := customerService.GetCustomer(customer.ID)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(exampleRes.CustomerResponse{Customer: data}, "获取成功", c)
	}
}

// GetCustomerList
// @Tags Customer
// @Summary 分页获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customerList [get]
func (e *CustomerApi) GetCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, customerList, total := customerService.GetCustomerInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     customerList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
