package middleware

import (
	"gin-starter/global"
	"gin-starter/model/system"
	"gin-starter/model/system/request"
	"gin-starter/service"
	"gin-starter/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
	"time"
)

var userService = service.AppService.SystemService.UserService

func ErrorToEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var username string
		if claims, ok := c.Get("claims"); ok {
			waitUse := claims.(*request.CustomClaims)
			username = waitUse.Username
		} else {
			id, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			err, user := userService.FindUserById(id)
			if err != nil {
				username = "Unknown"
			} else {
				username = user.Username
			}
		}
		body, _ := ioutil.ReadAll(c.Request.Body)
		record := system.Operations{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
		}
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		status := c.Writer.Status()
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		str := "接收到的请求为" + record.Body + "\n" + "请求方式为" + record.Method + "\n" + "报错信息如下" + record.ErrorMessage + "\n" + "耗时" + latency.String() + "\n"
		if status != 200 {
			subject := username + "" + record.Ip + "调用了" + record.Path + "报错了"
			if err := utils.ErrorToEmail(subject, str); err != nil {
				global.LOG.Error("ErrorToEmail Failed, err:", zap.Error(err))
			}
		}
	}
}
