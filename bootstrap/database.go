package bootstrap

import (
	"os"
	"time"

	"github.com/sjxiang/gohub/app/data/user"
	"github.com/sjxiang/gohub/pkg/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	
	var dbConfig gorm.Dialector	
	dsn := os.Getenv("MYSQL_DSN")
	
	dbConfig = mysql.New(mysql.Config{
		DSN: dsn,
	})


	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))  
	
	// 设置连接池（最大连接数、最大空闲连接数、每个连接的过期时间）
	database.SQLDB.SetMaxOpenConns(20)  // 打开
	database.SQLDB.SetMaxIdleConns(10)  // 空闲
	database.SQLDB.SetConnMaxLifetime(time.Duration(300) * time.Second)

	// 自动迁移
	database.DB.AutoMigrate((&user.User{}))
}