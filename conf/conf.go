package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var VerifyCodeTTL int


// 初始化配置项
func Init() {

	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
}

func IsLocal() bool {
	return os.Getenv("APP_ENV") == "local"
}


