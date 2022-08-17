// Package verifycode 用以发送手机验证码和邮箱验证码
package verifycode

import (
	"math/rand"
    "strings"
    "sync"
	"time"

	
	"github.com/sjxiang/gohub/config"
	"github.com/sjxiang/gohub/pkg/redis"
	"github.com/sjxiang/gohub/pkg/logger"
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
                RedisClient: redis.Redis,
                // 增加前缀保持数据库整洁，出问题调试时也方便
                KeyPrefix: config.Cfg.App.Name + ":verifycode:",
            },
        }
    })

    return internalVerifyCode
}


// CheckAnswer 检查用户提交的验证码是否正确，key 可以是手机号或者 Email
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {

    logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})

    // 方便开发，在非生产环境下，具备特殊前缀的手机号和 Email后缀，会直接验证成功
	// "verifycode.debug_email_suffix"
	// verifycode.debug_phone_prefix

    if config.Cfg.App.Islocal() && (strings.HasSuffix(key, "@test.com") || strings.HasPrefix(key, "000")) {
		return true
	}
	
    return vc.Store.Verify(key, answer, false)
}

// generateVerifyCode 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {

    // 生成随机码
    code := RandStringRunes(6)  // verifycode.code_length 6

    // 为方便开发，本地环境使用固定验证码
	if config.Cfg.App.Islocal() {
		code = "123456"  // verifycode.debug_code
	}

    logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

    // 将验证码及 KEY（邮箱或手机号）存放到 Redis 中并设置过期时间
    vc.Store.Set(key, code)
    return code
}



// RandStringRunes 返回随机字符串  生成长度为 length 随机数字字符串
func RandStringRunes(length int) string {
	var letterRunes = []rune("1234567890")  // abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
