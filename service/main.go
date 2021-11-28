package main

import (
	"fast-admin-service/core"
	"fast-admin-service/pkg/gorm"
	"fast-admin-service/pkg/log"
	"fast-admin-service/pkg/redis"
	"fast-admin-service/pkg/setting"
	"go.uber.org/zap"
)

func init() {
	log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE, "./log/gin-example.log")
	go func() {
		setting.Setup() // 初始化配置文件
		redis.Setup()   // redis初始化
		gorm.Setup()    // 初始化gorm
	}()
}

func main() {
	err := core.BootStarp().Run(":" + setting.ServerOptions.HttpPort)
	if err != nil {
		return
	}
}
