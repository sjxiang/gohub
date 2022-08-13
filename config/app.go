package config

type App struct {
	Name   		string `mapstructure:"name"`
	Port    	int    `mapstructure:"port"`   
	Debug       bool   `mapstructure:"debug"`  // 是否进入调试模式
	Key 		string `mapstructure:"key"`    // 加密会话，JWT 加密
	URL 		string `mapstructure:"app_url"`
	TimeZone 	string `mapstructure:"timezone"` // 设置时区，JWT 里会用到，日志记录里面也会用到
}

