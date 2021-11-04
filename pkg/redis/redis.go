package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

type RedisConnect struct {
	client *redis.Client
}

var (
	instance *RedisConnect
	once sync.Once
	ctx    = context.TODO()
)

func GetRedisConnInstance(address string) *RedisConnect {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr: address,
			Password: "",
			DB: 0,
		})
		if err := client.Ping(ctx).Err(); err != nil {
			return
		}
		instance=&RedisConnect{client: client}
	})
	return instance
}

func NewDatabase(address string) *redis.Client {
	return GetRedisConnInstance(address).client
}
