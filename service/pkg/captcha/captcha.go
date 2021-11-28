package captcha

import (
	"fast-admin-service/global"
	"fast-admin-service/pkg/cache"
	"github.com/mojocn/base64Captcha"
)

//var store = base64Captcha.DefaultMemStore
var store = cache.RedisStore{}

// GetCaptcha 获取验证码
func GetCaptcha() (string, string) {
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	c := base64Captcha.NewCaptcha(driver, store)

	// 获取
	id, b64s, err := c.Generate()
	if err != nil {
		global.ZAP_LOG.Error("Register GetCaptchaPhoto get base64Captcha has err:" + err.Error())
		return "", ""
	}
	return id, b64s
}

// Verify 验证验证码
func Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	// 同时在reids清理掉这个图片
	return store.Verify(id, val, true)
}
