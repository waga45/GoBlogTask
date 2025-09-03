package router

import (
	"GoBlogServer/internal/controller"
	"github.com/gin-gonic/gin"
)

// 初始化用户服务的路由
// 是否需要授权这里定义
func InitUserRouter(Router *gin.RouterGroup) {
	//公共不需要token授权就能访问
	userRoterPublic := Router.Group("user")
	//注册路由

	userRoterPublic.POST("/register", controller.Register)
	userRoterPublic.POST("/login", controller.Login)

	//需要授权路由
	userRoterPrivate := Router.Group("user")
	userRoterPrivate.GET("/info", controller.UserInfo)
}
