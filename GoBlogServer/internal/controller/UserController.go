package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(content *gin.Context) {
	fmt.Println("开始注册流程")
}

func Login(context *gin.Context) {
	fmt.Println("开始登入流程")
}

func UserInfo(context *gin.Context) {
	
}
