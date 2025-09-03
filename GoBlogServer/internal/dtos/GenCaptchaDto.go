package dtos

type GenCaptchaDto struct {
	CaptchaId string `json:"captchaId"`
	Image64   string `json:"image64"`
}
