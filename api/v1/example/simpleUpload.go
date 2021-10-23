package example

import (
	"gin-starter/global"
	"gin-starter/model/common/response"
	"gin-starter/model/example"
	"gin-starter/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SimpleUploadApi struct {
}

// SimpleUpload
//
// @Tags SimpleUpload
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "断点续传插件版示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"切片创建成功"}"
// @Router /SimpleUploadApi/upload [post]
func (s *SimpleUploadApi) SimpleUpload(c *gin.Context) {
	var chunk example.SimpleUploader
	_, header, _ := c.Request.FormFile("file")
	chunk.Filename = c.PostForm("filename")
	chunk.ChunkNumber = c.PostForm("chunkNumber")
	chunk.CurrentChunkSize = c.PostForm("currentChunkSize")
	chunk.Identifier = c.PostForm("identifier")
	chunk.TotalSize = c.PostForm("totalSize")
	chunk.TotalChunks = c.PostForm("totalChunks")
	var chunkDir = "./chunk/" + chunk.Identifier + "/"
	hasDir, _ := utils.PathExists(chunkDir)
	if !hasDir {
		if err := utils.CreateDir(chunkDir); err != nil {
			global.LOG.Error("创建目录失败!", zap.Error(err))
		}
	}
	chunkPath := chunkDir + chunk.Filename + chunk.ChunkNumber
	err := c.SaveUploadedFile(header, chunkPath)
	if err != nil {
		global.LOG.Error("切片创建失败!", zap.Error(err))
		response.FailWithMessage("切片创建失败", c)
		return
	}
	chunk.CurrentChunkPath = chunkPath
	err = simpleUploadService.SaveChunk(chunk)
	if err != nil {
		global.LOG.Error("切片创建失败!", zap.Error(err))
		response.FailWithMessage("切片创建失败", c)
		return
	} else {
		response.OkWithMessage("切片创建成功", c)
	}
}

// CheckFileMd5
// @Tags SimpleUpload
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @Produce  application/json
// @Param md5 query string true "md5"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SimpleUploadApi/checkFileMd5 [get]
func (s *SimpleUploadApi) CheckFileMd5(c *gin.Context) {
	md5 := c.Query("md5")
	err, chunks, isDone := simpleUploadService.CheckFileMd5(md5)
	if err != nil {
		global.LOG.Error("md5读取失败!", zap.Error(err))
		response.FailWithMessage("md5读取失败", c)
	} else {
		response.OkWithDetailed(gin.H{
			"chunks": chunks,
			"isDone": isDone,
		}, "查询成功", c)
	}
}

// MergeFileMd5
// @Tags SimpleUpload
// @Summary 合并文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param md5 query string true "md5"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"合并成功"}"
// @Router /SimpleUploadApi/mergeFileMd5 [get]
func (s *SimpleUploadApi) MergeFileMd5(c *gin.Context) {
	md5 := c.Query("md5")
	fileName := c.Query("fileName")
	err := simpleUploadService.MergeFileMd5(md5, fileName)
	if err != nil {
		global.LOG.Error("md5读取失败!", zap.Error(err))
		response.FailWithMessage("md5读取失败", c)
	} else {
		response.OkWithMessage("合并成功", c)
	}
}
