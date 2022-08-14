package bootstrap

import (
	"github.com/sjxiang/gohub/config"
	"github.com/sjxiang/gohub/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		config.Cfg.Log.FilePath, 
		config.Cfg.Log.MaxSize, 
		config.Cfg.Log.MaxBackup, 
		config.Cfg.Log.MaxAge, 
		config.Cfg.Log.Compress, 
		config.Cfg.Log.Type, 
		config.Cfg.Log.Level,
	)
}