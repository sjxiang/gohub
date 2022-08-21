// 实现 "gorm.io/gorm/logger" Interface，代替它，塞点私货进去

package logger

import (
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/sjxiang/gohub/pkg/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)


type GormLogger struct {
	ZapLogger 		*zap.Logger
	SlowThreshold 	time.Duration
}


// NewGormLogger，实例化一个 logger
// 使用示例：
// 		DB, err := gorm.Opn(dbConfig, &gorm.Config{
// 			Logger: logger.NewGormLogger()
// 		})
func NewGormLogger() GormLogger {
	return GormLogger{
		ZapLogger: Logger,             // 使用全局的  logger.Logger 对象
		SlowThreshold: 200 * time.Millisecond,  // 慢查询阈值，单位为 千分之一秒
	}
}


func (l GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return GormLogger{
		ZapLogger: l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
	}
}


func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Debugf(str, args...)
}


func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Warnf(str, args...)
}


func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Errorf(str, args...)
}


func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	
	// 获取运行时间
	elapsed := time.Since(begin)

	// 获取 SQL 请求和返回条数
	sql, rows := fc()

	// 通用字段
	logFields := []zap.Field{
		zap.String("sql", sql),
		zap.String("time", util.MicrosecondsStr(elapsed)),
		zap.Int64("rows", rows),
	}

	// Gorm 错误
	if err != nil {

		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.logger().Warn("Database ErrRecordNotFound", logFields...)
		} else {
			
			// 其它错误使用 error 等级
			logFields = append(logFields, zap.Error(err))
			l.logger().Error("Database Error", logFields...)
		}
	}

	// 慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.logger().Warn("Database Slow Log", logFields...)
	}

	// 记录所有 SQL 请求
	l.logger().Debug("'Database Query", logFields...)
}

// 辅助方法，确保 Zap 内置信息 Caller 的准确性（e.g. paginator/paginator.go:148）
func (l GormLogger) logger() *zap.Logger {
	
	// 跳过 gorm 内置调用
	var (
		gormPackage = filepath.Join("gorm.io", "gorm")  // gorm.io/gorm
		zapgormPackage = filepath.Join("moul.io", "zapgorm2")
	)

	// 减去 1 次封装，以及 1 次在 logger 初始化里添加 zap.AddCallerSkip(1)
	clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-2))

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			// 返回 1 个附带跳过行号的新的 zap logger
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}

	return l.ZapLogger
}
/*

2022/08/20 01:36:38 
/home/xsj/go/src/github.com/sjxiang/gohub/app/data/user/user_util.go:18
[0.336ms] [rows:1] 
SELECT count(*) FROM `users` WHERE phone = '18018001800'



2022-08-21 13:47:22	[35mDEBUG[0m	
user/user_util.go:18	
'Database Query	{"sql": "SELECT count(*) FROM `users` WHERE phone = '18018001800'", 
"time": "0.339 ms", "rows": 1}


还不如前者，除了写入日志这点

*/
