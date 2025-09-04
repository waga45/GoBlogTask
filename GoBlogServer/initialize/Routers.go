package initialize

import (
	"GoBlogServer/internal/container"
	"GoBlogServer/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// 初始化管理路由
func Routers(container *container.Container) *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.Default())
	// 或者使用自定义配置（生产环境推荐）
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:3000"},                                // 允许的源，谨慎使用 "*"
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},  // 允许的 HTTP 方法
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                                            // 允许客户端访问的响应头
		AllowCredentials: true,                                                                  // 是否允许发送 Cookie 等认证信息
		MaxAge:           12 * time.Hour,                                                        // 预检请求结果的缓存时间
	}))

	groupV1 := engine.Group("/v1")
	router.InitUserRouter(container, groupV1)
	router.InitPostsRouter(container, groupV1)
	return engine
}
