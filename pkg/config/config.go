// 负责配置信息
package config

import (
	"os"

	viperlib "github.com/spf13/viper" // 自定义包名，避免与内置 viper 实例冲突，类似 "_log"
)

// viper 库实例
var viper *viperlib.Viper


// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc



func init() {

	// 1. 初始化 Viper 库实例（就是上面那个全局变量）
	viper = viperlib.New()

	// 2. 配置类型
	viper.SetConfigType("env")

	// 3. 环境变量配置文件查找的路径，相对于 main.go
	viper.AddConfigPath(".")

	// 4. 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")

	// 5. 读取环境变量
	viper.AutomaticEnv()


	ConfigFuncs = make(map[string]ConfigFunc)
}


// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(env string) {

	// 1. 加载环境变量
	loadEnv(env)

	// 2. 注册配置信息
	loadConfig()

}


func loadConfig() {


}


func loadEnv(envSuffix string) {

	// 默认加载 .env 文件，如果有传参 --env=name，则加载 .env.name 文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filepath := ".env" + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			envPath = filepath
		}
	}

	// 加载 env（设置配置文件路径，读取配置）
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 监控 .env 文件，变更时重新加载
	viper.WatchConfig()

}

func Env() {

}

func Add() {

}



func Get