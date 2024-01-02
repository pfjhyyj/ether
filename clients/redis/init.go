package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"sync"
)

var (
	client redis.UniversalClient

	once sync.Once
)

func Init() {
	once.Do(func() {
		client = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    []string{viper.GetString("redis.host")},
			DB:       viper.GetInt("redis.database.biz"),
			Password: viper.GetString("redis.password"),
		})
	})
}

func GetRedisClient() redis.UniversalClient {
	Init()
	return client
}
