package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

var Cfg = Config{}

type Config struct {
	App           App           `mapstructure:"app"`
	Mysql         Mysql         `mapstructure:"mysql"`
	Redis         Redis         `mapstructure:"redis"`
}

/*

这里推荐使用 mapstructure 作为序列化标签

yaml 不支持 

	" AppSignExpire int64  `yaml:"app_sign_expire"` " 

这种下划线的标签

使用 mapstructure 值得注意的地方是，只要标签中使用了下划线等连接符，":"后就不能有空格。
比如： 
	AppSignExpire int64  `yaml:"app_sign_expire"`是可以被解析的
    AppSignExpire int64  `yaml: "app_sign_expire"` 不能被解析

*/

type App struct {
	Name   		string `mapstructure:"name"`
	Port    	int    `mapstructure:"port"`   
	Debug       bool   `mapstructure:"debug"`  // 是否进入调试模式
	Key 		string `mapstructure:"key"`    // 加密会话，JWT 加密
	URL 		string `mapstructure:"app_url"`
	TimeZone 	string `mapstructure:"timezone"` // 设置时区，JWT 里会用到，日志记录里面也会用到
}


type Mysql struct {
	DBName            string        `mapstructure:"dbname"`
	User              string        `mapstructure:"user"`
	Password          string        `mapstructure:"password"`
	Host              string        `mapstructure:"host"`
	MaxOpenConn       int           `mapstructure:"max_open_conn"`
	MaxIdleConn       int           `mapstructure:"max_idle_conn"`
	ConnMaxLifeSecond time.Duration `mapstructure:"conn_max_life_second"`
	TablePrefix       string        `mapstructure:"table_prefix"`
}


type Redis struct {
	Host        string `mapstructure:"host"`
	DB          int    `mapstructure:"db"`
	Password    string `mapstructure:"password"`
	MinIdleConn int    `mapstructure:"min_idle_conn"`
	PoolSize    int    `mapstructure:"pool_size"`
	MaxRetries  int    `mapstructure:"max_retries"`
}


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
