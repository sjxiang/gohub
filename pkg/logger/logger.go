// 处理日志相关逻辑

package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sjxiang/gohub/conf"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)


var Logger *zap.Logger

// InitLogger 日志初始化
func InitLogger(level string) {
	
	encoder := getEncoder()
	writer := getLogWriter()

	// 设置日志等级
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。")
	}

	// 初始化 core
	core := zapcore.NewCore(encoder, writer, logLevel)

	// 初始化 Logger
	Logger = zap.New(
		core, 
		zap.AddCaller(),   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),  // 封装了一层，调用文件去除一层（runtime.Caller(1)） 
		zap.AddStacktrace(zap.ErrorLevel),  // Error 才会显示 stacktrace
	)

	// 将自定义的 logger 替换为全局的 logger
	zap.ReplaceGlobals(Logger)
}


// getEncoder 设置 Log Entry 格式
func getEncoder() zapcore.Encoder {
	
	encoderConfig := zap.NewProductionEncoderConfig()  // 很多都是默认设置，不用再写一遍
	encoderConfig.EncodeTime = customTimeEncoder  // 时间戳格式

	// local dev 终端输出设置
	if conf.IsLocal() {
		
		// 终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		// 本地设置内置的 Console 解码器（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	
	return zapcore.NewJSONEncoder(encoderConfig)  // JSON 格式
}


// customTimeEncoder 自定义友好的时间戳格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}


// getLogWriter 多个输出流 
// 1. 日志文件 .log 
// 2. 控制台 os.Stdout  
func getLogWriter() zapcore.WriteSyncer {

	// 按照日期记录 log
	filename := "./tmp/logs/test.log"
	datename := time.Now().Format("2006-01-02")
	filename = strings.ReplaceAll(filename, "test", datename)

	// 滚动日志，使用 lumberjack 实现日志切割
	lumberJackLogger := &lumberjack.Logger{
		Filename: filename,  // 日志文件路径
		MaxSize: 64,         // 单个文件最大 64 M
		MaxBackups: 5,       // 多于 5 个文件后，清理较旧的日志；
		MaxAge: 30,  		 // 最多保存多少天，30 表示一月前的日志会被删除，0 表示不删
		Compress: false,     // 是否压缩
	}

	// 输出流配置
	if conf.IsLocal() {
		// 测试、开发环境 - 同时向控制台和日志文件输出 
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger)) 
	} else {
		// 生产环境只输出到文件中
		return zapcore.AddSync(lumberJackLogger)  // 同步写入 + 塞入句柄
	}
}

