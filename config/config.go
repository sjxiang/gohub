package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg = Config{}

type Config struct {
	App           App           `mapstructure:"app"`
	Mysql         Mysql         `mapstructure:"mysql"`
	Redis         Redis         `mapstructure:"redis"`
}

/* 

	*** 温馨小提示 tips

	这里推荐使用 mapstructure 作为序列化标签

	1. yaml 不支持 

		" AppSignExpire int64  `yaml:"app_sign_expire"` " 

		这种下划线的标签

	2. 使用 mapstructure 值得注意的地方是，只要标签中使用了下划线等连接符，":"后就不能有空格。
		比如： 
			AppSignExpire int64  `yaml:"app_sign_expire"`是可以被解析的
			AppSignExpire int64  `yaml: "app_sign_expire"` 不能被解析

*/


// 加载配置，失败直接 panic
func LoadConfig() {

	// 1. 初始化 viper 库的实例
	viper := viper.New()

	// 2. 设置配置文件路径
	viper.SetConfigFile("config/config.yml")

	// 3. 配置读取
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 4. 将配置映射成结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(err)
	}

	// 5. 监听配置文件变更，重新解析配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {  // 回调
		fmt.Println(e.Name)

		// Again，+1
		if err := viper.Unmarshal(&Cfg); err != nil {
			panic(err)
		}
	})
}
