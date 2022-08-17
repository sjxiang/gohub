package config


type VerifyCode struct {
	ExpireTime 		int64    `mapstructure:"expire_time"`     // 过期时间，单位 Minute
	TestExpireTime  int64    `mapstructure:"test_expire_time"`   // 本地开发测试模式下，延长过期时间，方便本地开发调试
}
