package vos

type UserLoginVO struct {
	Username string `json:"username"` //用户名称
	Password string `json:"password"` //用户密码
	Captcha  string `json:"captcha"`  //验证码
}
