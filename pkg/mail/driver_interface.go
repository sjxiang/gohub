package mail

// 不止一种实现呢？对吧！未雨绸缪，SMTP、POP3、IMAP ...
type Driver interface {
	// 检查验证码
	Send(email Email) bool
}