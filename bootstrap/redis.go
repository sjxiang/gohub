package bootstrap

import (
	"fmt"

	"github.com/sjxiang/gohub/config"
	"github.com/sjxiang/gohub/pkg/redis"
)



func SetupRedis() {
	
	addr := fmt.Sprintf("%v:%v", config.Cfg.Redis.Host, config.Cfg.Redis.Port)

	// 建立 Redis 连接
	redis.ConnectToRedis(
		addr,
		config.Cfg.Redis.Password,
		config.Cfg.Redis.Database,
	)

}