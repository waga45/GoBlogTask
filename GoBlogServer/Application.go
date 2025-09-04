package main

import (
	"GoBlogServer/config"
	"GoBlogServer/initialize" //全局路由管理
	"GoBlogServer/internal/component"
	"GoBlogServer/internal/container"
	"GoBlogServer/internal/repository/dbInit"
	"fmt"
	"go.uber.org/zap" //zap日志记录
	"os"
	"path/filepath"
)

func main() {
	//初始化日志
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	//config
	wd, _ := os.Getwd()
	appConfig, err := config.LoadConfig(filepath.Join(wd, "config/app_config.yaml"))
	if err != nil {
		fmt.Println(err)
		return
	}
	//token
	component.InitJwtManager(appConfig.Jwt.SecretKey, appConfig.Jwt.ExpireHour)
	//mysql
	db, err := dbInit.InitMysql(appConfig)
	if err != nil {
		zap.S().Error("数据库初始化链接失败...", err)
		return
	}
	//redis
	redisClient, _ := component.InitRedis(appConfig.Redis.Host, appConfig.Redis.Port, appConfig.Redis.Auth)
	defer redisClient.Close()
	//数据库和缓存展示全局管理
	Container := &container.Container{}
	Container.SetAppConfig(appConfig)
	Container.SetDbInstance(db)
	Container.SetRedisInstance(redisClient)
	//初始化路由+bean(这里bean按路由分组注入)
	engine := initialize.Routers(Container)
	//授权中间件
	//engine.Use(middleware.HandlerAuthWare())
	zap.S().Infof("GoBlog服务已启动...")
	err = engine.Run(fmt.Sprintf("%s:%d", appConfig.Server.Host, appConfig.Server.Port))
	if err != nil {
		zap.S().Panic("服务启动失败：", err)
	}
}
