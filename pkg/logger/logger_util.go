package logger

import (
	"encoding/json"
	
	"go.uber.org/zap"	
)


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

