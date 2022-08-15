// 处理日志相关逻辑

package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sjxiang/gohub/config"

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
	if config.Cfg.App.Islocal() {
		
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
	if config.Cfg.App.Islocal() {
		// 测试、开发环境 - 同时向控制台和日志文件输出 
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger)) 
	} else {
		// 生产环境只输出到文件中
		return zapcore.AddSync(lumberJackLogger)  // 同步写入 + 塞入句柄
	}
}


// ================================
// 日志辅助方法

func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error())) // struct 数据编码成 JSON
	}
	return string(b)
}


// Dump 调试专用，不会中断程序，会在终端打印出 warn 信息（高亮）。
// 第一个参数会使用 json.Marshal 进行渲染，第二个参数消息（可选）
// 				logger.Dump(user.User{Name:"test"})
//   			logger.Dump(user.User{Name:"test"}, "用户消息")

func Dump(value interface{}, msg ...string) {

	valueString := jsonString(value)

	// 判断第二个参数是否传参 msg
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("data", valueString))
	}
}


// LogIf() 减少代码中大量的 if err != nil { ... } 判断
func LogIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred", zap.Error(err))
	}
}

func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred", zap.Error(err))
	}
}

func LogInfof(err error) {
	if err != nil {
		Logger.Info("Error Occurred", zap.Error(err))
	}
}


// Debug 调试日志，详尽的程序日志
// 示例
// 			logger.Debug("Database", zap.String("sql", sql))
func Debug(moduleName string, fields ...zap.Field) {
	Logger.Debug(moduleName, fields...)
}

// Info 告知
func Info(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

// Warn 警告
func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

// Error 错误时记录，不应该中断程序，查看日志时重点关注
func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

// Fatal 级别同 Error()，写完 log 后调用 os.Exit(1) 退出程序
func Fatal(moduleName string, fields ...zap.Field) {
	Logger.Fatal(moduleName, fields...)
}



// DebugString 记录一条字符串类型的 debug 日志
// 示例
// 			logger.DebugString("SMS", "短信内容", String(result.RawResponse))

func DebugString(moduleName, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}


// DebugJSON 记录对象类型的 debug 日志，使用 json.Marshal 进行编码
// 示例
// 			logger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser{})

func DebugJSON(moduleName, name string, value interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName, name string, value interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

