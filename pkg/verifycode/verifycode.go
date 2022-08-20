// Package verifycode 用以发送手机验证码和邮箱验证码
package verifycode

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/sjxiang/gohub/conf"
	"github.com/sjxiang/gohub/pkg/cache"
	"github.com/sjxiang/gohub/pkg/logger"
	"github.com/sjxiang/gohub/pkg/mail"
	"github.com/sjxiang/gohub/pkg/util"
)

type VerifyCode struct {
    Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode


// NewVerifyCode 单例模式获取
func NewVerifyCode() *VerifyCode {
    once.Do(func() {
        internalVerifyCode = &VerifyCode{
            Store: &RedisStore{
                RedisClient: cache.Redis,
                // 增加前缀保持数据库整洁，出问题调试时也方便
                KeyPrefix: os.Getenv("APP_NAME") + ":Verifycode:",
            },
        }
    })

    return internalVerifyCode
}


func (vc *VerifyCode) SendEmail(email string) error {
    
    // 生成验证码
    code := vc.generateVerifyCode(email)

    // 方便本地和 API 自动测试（不用登录验证）
    if conf.IsLocal() && strings.HasSuffix(email, os.Getenv("VERIFYCODE_DEBUG_EMAIL_SUFFIX")) {
        return nil
    }

    // content := fmt.Sprintf("<h1>您的 Email 验证码是 %v </h1>", code)

    // 发送邮件
    mail.NewMailer().Send(mail.Email{
        From: os.Getenv("VERIFYCODE_FROM"),
        To: []string{os.Getenv("VERIFYCODE_TO")},
        Subject: "Gohub 登录验证",
        Text: []byte(fmt.Sprintf("您的 Email 验证码：%s", code)),
    })
    
    return nil
}


// CheckAnswer 检查用户提交的验证码是否正确，key 可以是手机号或者 Email
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {

    logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})

    // 方便开发，在非生产环境下，具备特殊前缀的手机号和 Email后缀，会直接验证成功
    if conf.IsLocal() && (strings.HasSuffix(key, os.Getenv("VERIFYCODE_DEBUG_EMAIL_SUFFIX")) || strings.HasPrefix(key, "VERIFYCODE_DEBUG_PHONE_PREFIX")) {
		return true
	}
	
    return vc.Store.Verify(key, answer, false)
}


// generateVerifyCode 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {

    // 生成随机码
    code := util.RandStringRunes(6)  // 验证码长度 6

    // 为方便开发，本地环境使用固定验证码
	if conf.IsLocal() {
		code = os.Getenv("VERIFYCODE_DEBUG_CODE") 
	}

    logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

    // 将验证码及 KEY（邮箱或者手机号码）存放到 Redis 中并设置过期时间
    vc.Store.Set(key, code)
    return code
}

