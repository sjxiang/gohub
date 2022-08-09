// Package routes 注册路由

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func RegisterApiRoutes(r *gin.Engine) {


	// v1 路由组
	v1 := r.Group("/v1")
	{
		// 注册路由
		v1.GET("/ping", func(ctx *gin.Context) {

			// 以 JSON 格式响应
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "pong",
			}) 
		})

	}
}



