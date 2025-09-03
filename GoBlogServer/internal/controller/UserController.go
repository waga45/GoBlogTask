package controller

import (
	"GoBlogServer/internal/constants"
	"GoBlogServer/internal/response"
	"GoBlogServer/internal/service"
	"GoBlogServer/internal/vos"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

// 注入服务
func Injection(s *service.UserService) *UserController {
	return &UserController{userService: s}
}

func (uc *UserController) Register(ctx *gin.Context) {
	fmt.Println("开始注册流程")
}

func (c *UserController) Login(ctx *gin.Context) {
	fmt.Println("开始登入流程")
	var vo vos.UserLoginVO
	err := ctx.ShouldBindJSON(&vo)
	if err != nil {
		response.Error(ctx, constants.ResponseError, "参数异常")
		return
	}

}

func UserInfo(context *gin.Context) {

}
