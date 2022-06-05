package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/samithiwat/samithiwat-backend/src/config"
)

func InitRedisConnect(conf *config.Redis) (cache *redis.Client, err error) {
	cache = redis.NewClient(&redis.Options{
		Addr: conf.Host,
		DB:   0,
	})

	return
}
