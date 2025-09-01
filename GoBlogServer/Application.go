package main

import (
	"GoBlogServer/initialize" //全局路由管理
	"go.uber.org/zap"         //zap日志记录
)

func main() {
	engine := initialize.Routers()
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infof("GoBlog服务已启动...")
	err := engine.Run(":8080")
	if err != nil {
		zap.S().Panic("服务启动失败：", err)
	}
}
