package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sjxiang/gohub/app/http/controllers/api/v1"
	"github.com/sjxiang/gohub/app/requests"
	"github.com/sjxiang/gohub/pkg/captcha"
	"github.com/sjxiang/gohub/pkg/logger"
	"github.com/sjxiang/gohub/pkg/response"
	"github.com/sjxiang/gohub/pkg/verifycode"
)

// VerifyCodeController 验证码控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}


// 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {

	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()

	// 记录错误日志（因为验证码是用户的入口，出错时应该记作 error 等级的日志）
	logger.LogIf(err)

	// 返回给用户
	response.JSON(c, gin.H{
		"captcha_id": id,
		"captcha_image": b64s,
	})

}


// SendUsingEmail 发送 Email 验证码
func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {
	
	// 1. 验证表单
	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	// 2. 发送邮件
	err := verifycode.NewVerifyCode().SendEmail(request.Email) 
	if err != nil {
		response.Abort500(c, "发送 Email 验证码错误")
	} else {
		response.Success(c)
	}
}