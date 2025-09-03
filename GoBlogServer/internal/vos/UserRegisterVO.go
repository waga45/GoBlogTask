package vos

type UserRegisterVO struct {
	Username  string `json:"username" validate:"required,min=2,max=32"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,max=32"`
	Code      string `json:"code" validate:"required"`
	CaptchaId string `json:"captchaId" validate:"required"`
}
