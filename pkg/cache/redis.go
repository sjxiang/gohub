// 客户端封装

package cache

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sjxiang/gohub/pkg/logger"
)

// RedisClient Redis 服务
type RedisClient struct {
	Client *redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// Redis 全局 Redis，使用 db 0
var Redis *RedisClient


// ConnectToRedis 连接 redis 数据库，设置全局的 Redis 对象
func ConnectToRedis(address string,  password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, password, db)
	})
}



func NewClient(address string, password string, db int) *RedisClient {

	// 初始化自定义的 RedisClient 实例
	rds := &RedisClient{}

	// 使用默认的 context
	rds.Context = context.Background()

	// 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr: address,
		Password: password,
		DB: db,
	})

	// 测试一下连接
	err := rds.Ping()
	logger.LogIf(err)

	return rds
}


// Ping 用以测试 redis 连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result();
	return err
}


// Set 存储 key 对应的 value，且设置 TTL 过期时间
func (rds RedisClient) Set(key string, value interface{}, TTL time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, TTL).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}

	return true
}


// Get 获取 key 对应的 value
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}

		return ""
	}

	return result
}


// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}

		return false
	}

	return true
}


// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}

	return true
}

// FlushDB 清空当前 redis db 里的所有数据
func (rds RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlashDB", err.Error())
		return false
	}

	return true
}

 
// Increment 当参数只有 1 个时，参数为 key，其值增加 1。
//           当参数有 2 个时，第一个参数为 key，第二个参数为要增加的值 int64 类型
func (rds RedisClient) Increment(params ...interface{}) bool {
	switch len(params) {
	case 1:
		key := params[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Incremet", err.Error())
			return false
		}
	case 2:
		key := params[0].(string)
		value := params[1].(int64)

		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Incremet", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Incremet", "参数过多")
		return false
	}

	return true
}



// Decrement 当参数只有 1 个时，参数为 key，其值减去 1。
//           当参数有 2 个时，第一个参数为 key，第二个参数为要减去的值 int64 类型
func (rds RedisClient) Decrement(params ...interface{}) bool {
	switch len(params) {
	case 1:
		key := params[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decremet", err.Error())
			return false
		}
	case 2:
		key := params[0].(string)
		value := params[1].(int64)

		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decremet", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decremet", "参数过多")
		return false
	}

	return true
}

