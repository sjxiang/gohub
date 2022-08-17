package mail

import (
	"sync"
)


type From struct {
	Addr string
	Name string
}

type Email struct {
	From 	From 
	To 		[]string
	Bcc 	[]string
	Cc 		[]string
	Subject string
	Text 	[]byte     // 明文信息（可选）
	HTML 	[]byte     // HTML 信息（可选）
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
	m := map[string]string{
		"Host": "localhost",
		"Port": "1025",
		"Username": "",
		"Password": "",
	}

	return mailer.Driver.Send(email, m)
}