package verifycode

import (
	"time"

	"github.com/sjxiang/gohub/conf"
	"github.com/sjxiang/gohub/pkg/cache"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *cache.RedisClient
	KeyPrefix string
}


func (s *RedisStore) Set(key string, value string) bool {

	TTL := time.Minute * time.Duration(30)  // 30 Min TTL 存活时间

	if conf.IsLocal() {
		TTL = time.Minute * time.Duration(300)  // 3 h
	}
	
	return s.RedisClient.Set(s.KeyPrefix+key, value, TTL)
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



