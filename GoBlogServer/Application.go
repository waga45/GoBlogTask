package main

import (
	"GoBlogServer/config"
	"GoBlogServer/initialize" //全局路由管理
	"GoBlogServer/internal/component"
	"GoBlogServer/internal/middleware"
	"fmt"
	"go.uber.org/zap" //zap日志记录
	"os"
	"path/filepath"
)

func main() {
	//初始化路由
	engine := initialize.Routers()
	//初始化日志
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infof("GoBlog服务已启动...")
	//加载配置
	wd, _ := os.Getwd()
	appConfig, err := config.LoadConfig(filepath.Join(wd, "config/app_config.yaml"))
	if err != nil {
		fmt.Println(err)
		return
	}
	//初始化redis单例
	redisClient, _ := component.InitRedis(appConfig.Redis.Host, appConfig.Redis.Port, appConfig.Redis.Auth)
	defer redisClient.Close()
	//token
	component.InitJwtManager(appConfig.Jwt.SecretKey, appConfig.Jwt.ExpireHour)
	//授权中间件
	engine.Use(middleware.HandlerAuthWare())
	err = engine.Run(fmt.Sprintf("%s:%d", appConfig.Server.Host, appConfig.Server.Port))
	if err != nil {
		zap.S().Panic("服务启动失败：", err)
	}
}
