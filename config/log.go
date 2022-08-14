package config

/*

	日志级别，必须是以下这些选项：
	"debug" - 信息量大，一般调试时打开。系统模块详细运行的日志，例如 HTTP 请求、数据库请求、发送邮件、发送短信
	"info" - 业务级别的运行日志，如用户登录、用户退出、订单撤销。
	"warn" - 感兴趣、需要引起关注的信息。 例如，调试时候打印调试信息（命令行输出会有高亮）。
	"error" - 记录错误信息。Panic 或者 Error。如数据库连接错误、HTTP 端口被占用等。一般生产环境使用的等级。
	
	以上级别从低到高，level 值设置的级别越高，记录到日志的信息就越少
	开发时推荐使用 "debug" 或者 "info" ，生产环境下使用 "error"

	日志的类型，可选：
	"single" 独立的文件
	"daily" 按照日期每日一个

*/

type Log struct {
	Level     string `mapstructure:"level"`
	Type      string `mapstructure:"type"`
	                                                /* --- 滚动日志配置 --- */ 
	FilePath  string `mapstructure:"file_path"`     // 日志文件路径
	MaxSize   int    `mapstructure:"max_size"`      // 每个日志文件保存的最大尺寸 单位：M
	MaxBackup int    `mapstructure:"max_backup"`    // 最多保存日志文件数，0 为不限，MaxAge 到了还是会删
	MaxAge    int    `mapstructure:"max_age"`       // 最多保存多少天，7 表示一周前的日志会被删除，0 表示不删
	Compress  bool   `mapstructure:"compress"`      // 是否压缩，压缩日志不方便查看，设置为 false（压缩可节省空间）
} 


