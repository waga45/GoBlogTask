package router

import (
	"GoBlogServer/internal/container"
	"GoBlogServer/internal/controller"
	"GoBlogServer/internal/middleware"
	"GoBlogServer/internal/repository/mapper"
	"GoBlogServer/internal/service"
	"github.com/gin-gonic/gin"
)

// 初始化用户服务的路由
// 是否需要授权这里定义
func InitUserRouter(Container *container.Container, Router *gin.RouterGroup) {
	//InitBean
	userMapper := mapper.NewUserMapper(Container.GetGormDb())
	userService := service.NewUserService(userMapper)
	captchaservice := service.NewCaptchaService()
	userController := controller.NewUserController(userService, captchaservice)

	//公共不需要token授权就能访问
	userRoterPublic := Router.Group("user")
	//注册路由
	userRoterPublic.GET("/captcha", userController.Captcha)
	userRoterPublic.POST("/register", userController.Register)
	userRoterPublic.POST("/login", userController.Login)

	//需要授权路由
	userRoterPrivate := Router.Group("user")
	//需要授权
	userRoterPrivate.Use(middleware.HandlerAuthWare())
	userRoterPrivate.GET("/info", userController.UserInfo)

}
