package config

import (
	"time"
)


type Mysql struct {
	DBName            string        `mapstructure:"dbname"`  // 数据库连接信息
	Username          string        `mapstructure:"username"`
	Password          string        `mapstructure:"password"`
	Host              string        `mapstructure:"host"`
	Port              int           `mapstructure:"port"`
	Charset           string        `mapstructure:"charset"`
	MaxOpenConn       int           `mapstructure:"max_open_conn"`  // 连接池配置
	MaxIdleConn       int           `mapstructure:"max_idle_conn"`
	ConnMaxLifeSecond time.Duration `mapstructure:"conn_max_life_second"`
	TablePrefix       string        `mapstructure:"table_prefix"`
}
