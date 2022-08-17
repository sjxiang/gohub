package config


type Captcha struct {
	Height 			int      `mapstructure:"height"`               // 验证码图片高度
	Width 			int      `mapstructure:"width"`                // 验证码图片宽度
	Length 			int		 `mapstructure:"length"`               // 验证码的长度
	Maxskew 		float64  `mapstructure:"maxskew"`              // 数字的最大倾斜角度
	Dotcount 		int      `mapstructure:"dotcount"`             // 图片背景里的混淆点数量
	ExpireTime 		int64    `mapstructure:"expire_time"`          // 过期时间，单位 Minute
	DebugExpireTime int64    `mapstructure:"debug_expire_time"`    // debug 模式下的过期时间，方便本地开发调试
	Testingkey 		string   `mapstructure:"testing_key"`          // 非生产环境，使用此 key 可跳过验证，方便测试
}
