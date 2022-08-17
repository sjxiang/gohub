package verifycode

import (
	"time"

	"github.com/sjxiang/gohub/config"
	"github.com/sjxiang/gohub/pkg/redis"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix string
}

func (s *RedisStore) Set(key string, value string) bool {

	ExpiredTime := time.Minute * time.Duration(config.Cfg.VerifyCode.ExpireTime)  // 15 Min
	
	// local dev 本地开发环境
	if config.Cfg.App.Islocal() {
		ExpiredTime = time.Minute * time.Duration(config.Cfg.VerifyCode.TestExpireTime)  // 3 h
	}

	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpiredTime)
}

func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}

	return val
}

func (s *RedisStore)Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}



