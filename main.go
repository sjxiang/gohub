package main

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
	"github.com/sjxiang/gohub/bootstrap"
	"github.com/sjxiang/gohub/config"
)



func init() {
	
	// 加载 config 目录下的配置信息 
	config.LoadConfig()

	// 初始化 DB
	bootstrap.SetupDB()

	// 初始化 Logger
	bootstrap.SetupLogger()
	
}


func main() {
	
	// new 一个 Gin Engine 实例（指针对象，不会被逃逸分析或垃圾回收干掉，尽情配置）
	router := gin.New() 

	// 初始化路由绑定
	bootstrap.SetupRoute(router)


	// 运行服务，指定监听端口为 3000
	err := router.Run(fmt.Sprintf(":%v", config.Cfg.App.Port))
	if err != nil {

		// 错误处理，端口被占用或其他错误
		fmt.Println(err.Error())
	}

}

