package router

import (
	"GoBlogServer/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRoter := Router.Group("user")
	//注册路由
	userRoter.POST("/register", controller.Register)
}
