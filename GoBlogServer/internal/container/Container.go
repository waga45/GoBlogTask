package container

import (
	"GoBlogServer/config"
	"GoBlogServer/internal/component"
	"gorm.io/gorm"
)

// 全局容器
type Container struct {
	appConfig   *config.AppConfig
	db          *gorm.DB
	redisClient *component.RedisClient
}

func (c *Container) GetGormDb() *gorm.DB {
	return c.db
}

func (c *Container) GetRedisClient() *component.RedisClient {
	return c.redisClient
}

func (c *Container) GetAppConfig() *config.AppConfig {
	return c.appConfig
}
func (c *Container) SetAppConfig(appConfig *config.AppConfig) {
	c.appConfig = appConfig
}
func (c *Container) SetDbInstance(_db *gorm.DB) {
	c.db = _db
}

func (c *Container) SetRedisInstance(rds *component.RedisClient) {
	c.redisClient = rds
}
