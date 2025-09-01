package initialize

import (
	"GoBlogServer/router"
	"github.com/gin-gonic/gin"
)

// 初始化管理路由
func Routers() *gin.Engine {
	engine := gin.Default()
	groupV1 := engine.Group("/v1")
	router.InitUserRouter(groupV1)
	return engine
}
