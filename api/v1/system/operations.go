package system

import (
	"gin-starter/global"
	"gin-starter/model/common/request"
	"gin-starter/model/common/response"
	"gin-starter/model/system"
	systemReq "gin-starter/model/system/request"
	"gin-starter/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Operations struct {
}

// CreateOperations
//
// @Tags Operations
// @Summary 创建Operations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Operations true "创建Operations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /operations/createOperations [post]
func (s *Operations) CreateOperations(c *gin.Context) {
	var operations system.Operations
	_ = c.ShouldBindJSON(&operations)
	if err := operationsService.CreateOperations(operations); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOperations
//
// @Tags Operations
// @Summary 删除Operations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Operations true "Operations模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /operations/deleteOperations [delete]
func (s *Operations) DeleteOperations(c *gin.Context) {
	var operations system.Operations
	_ = c.ShouldBindJSON(&operations)
	if err := operationsService.DeleteOperations(operations); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOperationsByIds
// @Tags Operations
// @Summary 批量删除Operations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Operations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /operations/deleteOperationsByIds [delete]
func (s *Operations) DeleteOperationsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := operationsService.DeleteOperationsByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindOperations
//
// @Tags Operations
// @Summary 用id查询Operations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Operations true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /operations/findOperations [get]
func (s *Operations) FindOperations(c *gin.Context) {
	var operations system.Operations
	_ = c.ShouldBindQuery(&operations)
	if err := utils.Verify(operations, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, reOperations := operationsService.GetOperations(operations.ID); err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"reOperations": reOperations}, "查询成功", c)
	}
}

// GetOperationsList
//
// @Tags Operations
// @Summary 分页获取Operations列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.OperationsSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /operations/getOperationsList [get]
func (s *Operations) GetOperationsList(c *gin.Context) {
	var pageInfo systemReq.OperationsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := operationsService.GetOperationsInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
