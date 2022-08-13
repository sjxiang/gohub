package bootstrap

import (
	"fmt"
	"time"

	"github.com/sjxiang/gohub/config"
	"github.com/sjxiang/gohub/pkg/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)



// SetupDB 初始化数据库和 ORM
func SetupDB() {
	
	var dbConfig gorm.Dialector	
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		config.Cfg.Mysql.Username,
		config.Cfg.Mysql.Password,
		config.Cfg.Mysql.Host,
		config.Cfg.Mysql.Port,
		config.Cfg.Mysql.DBName,
		config.Cfg.Mysql.Charset,
	)

	dbConfig = mysql.New(mysql.Config{
		DSN: dsn,
	})


	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	
	// 设置连接池（最大连接数、最大空闲连接数、每个连接的过期时间）
	database.SQLDB.SetMaxOpenConns(config.Cfg.Mysql.MaxOpenConn)
	database.SQLDB.SetMaxIdleConns(config.Cfg.Mysql.MaxIdleConn)
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.Cfg.Mysql.ConnMaxLifeSecond) * time.Second)
}