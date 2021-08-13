package example

import (
	"gin-starter/global"
	"gin-starter/model/common/request"
	"gin-starter/model/common/response"
	"gin-starter/model/example"
	exampleRes "gin-starter/model/example/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileTransferApi struct {
}

// UploadFile
// @Tags FileTransfer
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileTransfer/upload [post]
func (u *FileTransferApi) UploadFile(c *gin.Context) {
	var file example.FileTransfer
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file = fileTransferService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaFileResponse{File: file}, "上传成功", c)
}

// DeleteFile
// @Tags FileTransfer
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body example.FileTransfer true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileTransfer/deleteFile [post]
func (u *FileTransferApi) DeleteFile(c *gin.Context) {
	var file example.FileTransfer
	_ = c.ShouldBindJSON(&file)
	if err := fileTransferService.DeleteFile(file); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetFileList
// @Tags FileTransfer
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileTransfer/getFileList [post]
func (u *FileTransferApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := fileTransferService.GetFileRecordInfoList(pageInfo)
	if err != nil {
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
