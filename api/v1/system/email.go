package system

import (
	"gin-starter/global"
	"gin-starter/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// EmailTest
//
// @Tags System
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/emailTest [post]
func (s *Systems) EmailTest(c *gin.Context) {
	if err := emailService.EmailTest(); err != nil {
		global.LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
