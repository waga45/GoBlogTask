package controller

import (
	"GoBlogServer/internal/constants"
	"GoBlogServer/internal/dtos"
	"GoBlogServer/internal/response"
	"GoBlogServer/internal/service"
	"GoBlogServer/internal/vos"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService    *service.UserService
	captchaService *service.CaptchaService
}

// 注入服务
func NewUserController(s *service.UserService, c *service.CaptchaService) *UserController {
	return &UserController{userService: s, captchaService: c}
}

/*
*
生成验证码
*/
func (uc *UserController) Captcha(ctx *gin.Context) {
	id, bs64img, err := uc.captchaService.GenCaptcha(ctx)
	if err != nil {
		response.Error(ctx, constants.ResponseError, "验证码生成失败")
		return
	}
	var captchaDto = dtos.GenCaptchaDto{CaptchaId: id, Image64: *bs64img}
	response.Success(ctx, captchaDto)
}

/*
**
注册API
*/
func (uc *UserController) Register(ctx *gin.Context) {
	fmt.Println("开始注册流程")
	var vo vos.UserRegisterVO
	err := ctx.ShouldBindJSON(&vo)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
		return
	}
	vSuccess := uc.captchaService.VerifyCaptcha(vo.CaptchaId, vo.Code)
	if !vSuccess {
		response.Error(ctx, constants.ResponseError, "验证码错误！")
		return
	}
	result, err := uc.userService.RegisterUser(&vo)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
		return
	}

	response.Success(ctx, result)
}

/*
**
登入api
*/
func (uc *UserController) Login(ctx *gin.Context) {
	fmt.Println("开始登入流程")
	var vo vos.UserLoginVO
	err := ctx.ShouldBindJSON(&vo)
	if err != nil {
		response.Error(ctx, constants.ResponseError, "参数异常")
		return
	}
	vSuccess := uc.captchaService.VerifyCaptcha(vo.CaptchaId, vo.Code)
	if !vSuccess {
		response.Error(ctx, constants.ResponseError, "验证码错误！")
		return
	}
	user, err := uc.userService.UserLogin(&vo)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
	}
	response.Success(ctx, user)
}

func (uc *UserController) UserInfo(context *gin.Context) {

}
