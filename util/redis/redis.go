package redis

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	once sync.Once
	rdb  *redis.Client
)

func NewRedis() *redis.Client {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})
	})

	return rdb
}
