package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"sync"
)

type CaptchaService struct {
}

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{}
}

var captchaStore = sync.Map{}

func (c *CaptchaService) GenCaptcha(ctx *gin.Context) (captchaId string, image64 *string, err error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.5, 8)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return "", nil, err
	}
	zap.S().Infof("验证码Id:", id, " 码：", answer)
	captchaStore.Store(id, answer)
	return id, &b64s, nil
}

/*
*
验证
*/
func (c *CaptchaService) VerifyCaptcha(captchaId string, code string) bool {
	if len(captchaId) <= 0 || len(code) <= 0 {
		return false
	}
	answer, b := captchaStore.Load(captchaId)

	if !b {
		return false
	}
	if code != answer.(string) {
		return false
	}
	captchaStore.Delete(captchaId)
	return true
}
