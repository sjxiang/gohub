// 处理日志相关逻辑

package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sjxiang/gohub/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger，全局 Logger 对象
var Logger *zap.Logger

// InitLogger 日志初始化
func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {
	
	// 获取日志写入介质
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)

	// 设置日志等级，具体请见 config/log.go 文件
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}

	// 初始化 core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	// 初始化 Logger
	Logger = zap.New(
		core, 
		zap.AddCaller(), 
		zap.AddCallerSkip(1), 
		zap.AddStacktrace(zap.ErrorLevel),
	)

	// 将自定义的 logger 替换为全局的 logger
	zap.ReplaceGlobals(Logger)
}


// getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {
	
	// 日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		
		// 日志记录前缀设置
		TimeKey: "时间戳",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "caller",  // 代码调用
		FunctionKey: zapcore.OmitKey,
		MessageKey: "msg",
		StacktraceKey: "stacktrace",

		LineEnding: zapcore.DefaultLineEnding,  // 每行日志的结尾添加 "\n"
		EncodeLevel: zapcore.CapitalLevelEncoder,  // 日志级别名称大写，如 ERROR、INFO
		EncodeTime: customTimeEncoder,  // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,  // 执行时间，以 '秒' 为单位
		EncodeCaller: zapcore.ShortCallerEncoder,  // Caller 短格式，如 types/converter.go:17；长格式为绝对路径，~/go/src/。
	} 

	// 本地环境配置
	if config.Cfg.App.Islocal() {
		
		// 终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		// 本地设置内置的 Console 解码器（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}


// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}


// getLogWriter 日志记录介质。Gohub 使用了两种介质 os.Stdout 和 log 文件
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer{

	// 如果配置了按照日期记录日志文件
	if logType == "daily" {
		logname := time.Now().Format("2006-01-02")
		filename = strings.ReplaceAll(filename, "logs.log", logname)
	}

	// 滚动日志，详见 config/log.go
	lumberJackLogger := &lumberjack.Logger{
		Filename: filename,
		MaxSize: maxSize,
		MaxBackups: maxBackup,
		MaxAge: maxAge,
		Compress: compress,
	}

	// 配置输出介质
	if config.Cfg.App.Islocal() {
		// 本地开发终端打印和记录文件、
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		// 生产环境只记录文件
		return zapcore.AddSync(lumberJackLogger)
	}


}
