package vos

type UserLoginVO struct {
	Username  string `json:"username" validate:"required"` //用户名称
	Password  string `json:"password" validate:"required"` //用户密码
	Code      string `json:"code" validate:"required"`
	CaptchaId string `json:"captchaId" validate:"required"`
}
