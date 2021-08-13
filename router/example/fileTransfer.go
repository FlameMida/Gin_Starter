package example

import (
	"gin-starter/api/v1"
	"github.com/gin-gonic/gin"
)

type FileTransferRouter struct {
}

func (e *FileTransferRouter) InitFileTransferRouter(Router *gin.RouterGroup) {
	fileTransferRouter := Router.Group("fileTransfer")
	var exaFileTransferApi = v1.AppApi.ExampleApi.FileTransferApi
	{
		fileTransferRouter.POST("/upload", exaFileTransferApi.UploadFile)                                 // 上传文件
		fileTransferRouter.POST("/getFileList", exaFileTransferApi.GetFileList)                           // 获取上传文件列表
		fileTransferRouter.POST("/deleteFile", exaFileTransferApi.DeleteFile)                             // 删除指定文件
		fileTransferRouter.POST("/breakpointContinue", exaFileTransferApi.BreakpointContinue)             // 断点续传
		fileTransferRouter.GET("/findFile", exaFileTransferApi.FindFile)                                  // 查询当前文件成功的切片
		fileTransferRouter.POST("/breakpointContinueFinish", exaFileTransferApi.BreakpointContinueFinish) // 查询当前文件成功的切片
		fileTransferRouter.POST("/removeChunk", exaFileTransferApi.RemoveChunk)                           // 查询当前文件成功的切片
	}
}
