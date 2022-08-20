// Package routes 注册路由

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/gohub/app/http/controllers/api/v1/auth"
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

		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)

			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
		}

	}
}



