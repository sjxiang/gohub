// 处理图片验证码逻辑

package captcha

import (
	"fmt"
	"sync"

	"github.com/mojocn/base64Captcha"
	"github.com/sjxiang/gohub/config"
	"github.com/sjxiang/gohub/pkg/redis"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}


// once 确保 internalCaptcha 对象值初始化一次
var once sync.Once

// internalCaptcha 内部使用的 Captcha 对象
var internalCaptcha *Captcha

// 
func NewCaptcha() *Captcha {
	once.Do(func() {

		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}

		// 使用全局 Redis 对象，并配置存储 key 的前缀
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix: fmt.Sprintf("%s :Captcha", config.Cfg.App.Name),
		}

		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			config.Cfg.Captcha.Height,
			config.Cfg.Captcha.Width,
			config.Cfg.Captcha.Length,
			config.Cfg.Captcha.Maxskew,
			config.Cfg.Captcha.Dotcount,
		)

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store) 
	})

	return internalCaptcha
}

// GenerateCaptcha 生成图片验证码
func (c *Captcha) GenerateCaptcha(id string, b64s string, err error) {
	c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证码是否正确
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {

	if id == config.Cfg.Captcha.Testingkey {
		return true
	}

	// 第 3 个参数，是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单错误提交需要多次输入图片验证码
	return c.Base64Captcha.Verify(id, answer, false)
}
