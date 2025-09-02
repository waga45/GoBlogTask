package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseWriter struct {
	C *gin.Context
}

func NewResponseWriter(c *gin.Context) *ResponseWriter {
	return &ResponseWriter{C: c}
}

// JSON 返回JSON响应
func (rw *ResponseWriter) JSON(response *BaseResponse) {
	rw.C.JSON(http.StatusOK, response)
}

// Success 返回成功响应
func (rw *ResponseWriter) Success(data interface{}) {
	rw.JSON(BuildSuccess(data))
}

// SuccessWithMessage 返回带消息的成功响应
func (rw *ResponseWriter) SuccessWithMessage(message string, data interface{}) {
	rw.JSON(BuildSuccessWithMessage(message, data))
}

// Error 返回错误响应
func (rw *ResponseWriter) Error(code int, message string) {
	rw.JSON(BuildError(code, message))
}

// PredefinedError 返回预定义错误
func (rw *ResponseWriter) PredefinedError(response *BaseResponse) {
	rw.JSON(response)
}

// 简化函数，无需创建ResponseWriter实例

// JSON 直接返回JSON响应
func JSON(c *gin.Context, response *BaseResponse) {
	c.JSON(http.StatusOK, response)
}

// Success 直接返回成功响应
func Success(c *gin.Context, data interface{}) {
	JSON(c, BuildSuccess(data))
}

// SuccessWithMessage 直接返回带消息的成功响应
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	JSON(c, BuildSuccessWithMessage(message, data))
}

// Error 直接返回错误响应
func Error(c *gin.Context, code int, message string) {
	JSON(c, BuildError(code, message))
}

// PredefinedError 直接返回预定义错误
func PredefinedError(c *gin.Context, response *BaseResponse) {
	JSON(c, response)
}
