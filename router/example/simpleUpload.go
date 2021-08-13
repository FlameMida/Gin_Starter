package example

import (
	"gin-starter/api/v1"
	"github.com/gin-gonic/gin"
)

type SimpleUploaderRouter struct{}

func (e *SimpleUploaderRouter) InitSimpleUploaderRouter(Router *gin.RouterGroup) {
	simpleUploadRouter := Router.Group("simpleUpload")
	var simpleUploadApi = v1.AppApi.ExampleApi.SimpleUploadApi
	{
		simpleUploadRouter.POST("upload", simpleUploadApi.SimpleUpload)      // 上传功能
		simpleUploadRouter.GET("checkFileMd5", simpleUploadApi.CheckFileMd5) // 文件完整度验证
		simpleUploadRouter.GET("mergeFileMd5", simpleUploadApi.MergeFileMd5) // 合并文件
	}
}
