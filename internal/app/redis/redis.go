package redis

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
)

var Module = fx.Module("redis",
	fx.Provide(
		NewRedisClient,
	),
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
