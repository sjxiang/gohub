package mail

import (
	"net/smtp"
	"os"

	emailPKG "github.com/jordan-wright/email"
	"github.com/sjxiang/gohub/pkg/logger"
)

// SMTP 实现 email.Driver interface
type SMTP struct{}

// Send 实现 email.Driver interface 的 Send 方法
func (S *SMTP) Send(email Email) bool {
	
	e := emailPKG.NewEmail()

	e.From = email.From
	e.To = email.To
	e.Subject = email.Subject
	e.Text = email.Text

	logger.DebugJSON("发送邮件", "发送详情", e)

	err := e.Send(
		"smtp.qq.com",
		smtp.PlainAuth(
			"",
			os.Getenv("VERIFYCODE_FROM"),  // 服务器邮箱账号
			os.Getenv("VERIFYCODE_QQEmailAuthCode"),  // 授权码
			"smtp.qq.com",
		),
	)

	if err != nil {
		logger.ErrorString("发送邮件", "发件出错", err.Error())
		return false
	}

	logger.DebugString("发送邮件", "发件成功", "")
	return true
}