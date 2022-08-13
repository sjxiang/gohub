package config


type Redis struct {
	Host        string `mapstructure:"host"`
	DB          int    `mapstructure:"db"`
	Password    string `mapstructure:"password"`
	MinIdleConn int    `mapstructure:"min_idle_conn"`
	PoolSize    int    `mapstructure:"pool_size"`
	MaxRetries  int    `mapstructure:"max_retries"`
}
