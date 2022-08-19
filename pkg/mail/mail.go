package mail

import (
	"sync"
)


type Email struct {
	From 	string      // 发送方的邮箱，服务器邮箱
	To 		[]string    // 接收方的邮箱，注册登录用户的真实邮箱
	Subject string      // 邮件主题 "Gohub 登录"
	Text 	[]byte      // 邮件内容 "验证码 123456"（明文信息） 
}


type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer


// NewMailer 单例模式获取
func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})

	return internalMailer
}


func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email)
}