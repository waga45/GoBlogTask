package response

import (
	"GoBlogServer/internal/constants"
	"time"
)

type BaseResponse struct {
	Code      int         `json:"code"`                 // 状态码
	Message   string      `json:"message"`              // 消息
	Data      interface{} `json:"data,omitempty"`       // 数据
	Timestamp int64       `json:"timestamp"`            // 时间戳
	RequestID string      `json:"request_id,omitempty"` // 请求ID
}

func NewResponse(code int, message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:      code,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
}

/*
*
成功
*/
func BuildSuccess(data interface{}) *BaseResponse {
	return NewResponse(constants.ResponseSuccess, "success", data)
}

/*
**
带消息的成功响应
*/
func BuildSuccessWithMessage(message string, data interface{}) *BaseResponse {
	return NewResponse(constants.ResponseSuccess, message, data)
}

/*
*
通用错误返回
*/
func BuildError(code int, message string) *BaseResponse {
	return NewResponse(code, message, nil)
}
