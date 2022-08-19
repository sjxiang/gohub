package bootstrap

import (
	"os"
	"strconv"

	"github.com/sjxiang/gohub/pkg/cache"
)



func SetupRedis() {

	db, _ := strconv.Atoi(os.Getenv("REDIS_DB")) 

	// 建立 Redis 连接
	cache.ConnectToRedis(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PW"),
		db,
	)
}