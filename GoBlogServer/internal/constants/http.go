package constants

// 请求头字段
const (
	AuthorizationKey = "Authorization"
	UserIdKey        = "userId"
)

// HTTP状态码相关常量
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusInternalServerError = 500
)
