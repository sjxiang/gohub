package bootstrap

import (
	"github.com/sjxiang/gohub/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	
	var level = "debug"

	logger.InitLogger(level)
}


/*

日志级别，必须是以下这些选项：

	"debug" - 信息量大，一般调试时打开。系统模块详细运行的日志，例如 HTTP 请求、数据库请求、发送邮件、发送短信
	"info" - 业务级别的运行日志，如用户登录、用户退出、订单撤销。
	"warn" - 感兴趣、需要引起关注的信息。 例如，调试时候打印调试信息（命令行输出会有高亮）。
	"error" - 记录错误信息。Panic 或者 Error。如数据库连接错误、HTTP 端口被占用等。一般生产环境使用的等级。
	
	以上级别从低到高，level 值设置的级别越高，记录到日志的信息就越少
	开发时推荐使用 "debug" 或者 "info" ，生产环境下使用 "error"

*/