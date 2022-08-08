package main

import (

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	
	// new 一个 Gin Engine 实例
	r := gin.New() 

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())


	// 注册路由
	r.GET("/ping", func(ctx *gin.Context) {

		// 以 JSON 格式响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		}) 
	})


	// 处理 404 请求


	
	// 运行服务，指定端口为 8080
	r.Run(":8080")
}

