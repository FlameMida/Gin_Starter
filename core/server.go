package core

import (
	"fmt"
	"gin-starter/global"
	"gin-starter/initialize"
	"gin-starter/service/system"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	// 从db加载jwt数据
	if global.DB != nil {
		system.LoadAll()
	}
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf("[GIN-STARTER]文档地址:http://127.0.0.1%s/swagger/index.html \n", address)
	global.LOG.Error(s.ListenAndServe().Error())

}
