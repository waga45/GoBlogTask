package middleware

import (
	"GoBlogServer/internal/component"
	"GoBlogServer/internal/constants"
	"GoBlogServer/internal/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func HandlerAuthWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//验证token
		token := c.GetHeader(constants.AuthorizationKey)
		if len(token) == 0 {
			response.Error(c, constants.ResponseAuthFail, "授权失败")
			return
		}
		payload, err := component.GetJwtInstance().ParserToken(token)
		if err != nil {
			zap.S().Error("Token解析校验失败", err)
			response.Error(c, constants.ResponseAuthFail, "授权失败")
			return
		}
		//解析注入用户id
		c.Header(constants.UserIdKey, string(strconv.Itoa(payload.UserId)))
		c.Next()
	}
}
