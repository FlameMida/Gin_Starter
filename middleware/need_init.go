package middleware

import (
	"gin-starter/global"
	"gin-starter/model/common/response"
	"github.com/gin-gonic/gin"
)

// NeedInit 处理跨域请求,支持options访问
func NeedInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.DB == nil {
			response.OkWithDetailed(gin.H{
				"needInit": true,
			}, "前往初始化数据库", c)
			c.Abort()
		} else {
			c.Next()
		}
		// 处理请求
	}
}
