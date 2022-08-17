package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/gohub/bootstrap"
	"github.com/sjxiang/gohub/config"
	"github.com/sjxiang/gohub/pkg/captcha"
	"github.com/sjxiang/gohub/pkg/logger"
)



func init() {
	
	// 加载 config 目录下的配置信息 
	config.LoadConfig()

	// 初始化 DB
	bootstrap.SetupDB()

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 初始化 Redis
	bootstrap.SetupRedis()
	
}


func main() {
	
	// 设置 gin 的运行模式，支持 debug、release、test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	gin.SetMode(gin.ReleaseMode)
	
	// new 一个 Gin Engine 实例（指针对象，不会被逃逸分析或垃圾回收干掉，尽情配置）
	router := gin.New() 

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	logger.Dump(captcha.NewCaptcha().VerifyCaptcha("aJaNjbG7J6MFvJVsDY1G", "568092"), "正确的答案")
	logger.Dump(captcha.NewCaptcha().VerifyCaptcha("aJaNjbG7J6MFvJVsDY1G", "568392"), "错误的答案")


	// 运行服务，指定监听端口为 3000
	err := router.Run(fmt.Sprintf(":%v", config.Cfg.App.Port))
	if err != nil {

		// 错误处理，端口被占用或其他错误
		fmt.Println(err.Error())
	}

}

