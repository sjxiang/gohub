// Package bootstrap 处理程序初始化逻辑

package bootstrap

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/gohub/routes"
)


// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {
	
	// 注册中间件
	registerMiddleWare(router)

	// 注册 API 路由
	routes.RegisterApiRoutes(router)

	// 配置 404 路由
	setupNoFoundHandler(router)
	
}
  


func registerMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}



func setupNoFoundHandler(router *gin.Engine) {


	// 处理 404 请求
	router.NoRoute(func(ctx *gin.Context) {

		// 获取 header 信息的 "Accept" 信息
		acceptStr := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptStr, "text/html") {

			// 如果是 HTML 的话
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {

			// 默认返回 JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确",
			})	
		}

	})
}