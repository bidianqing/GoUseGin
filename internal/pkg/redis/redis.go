package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx = context.Background()

func NewRedisClient(configuration string) {
	opts, err := redis.ParseURL(configuration)
	if err != nil {
		panic(err)
	}
	client = redis.NewClient(opts)
}

func CloseRedisClient() {
	client.Close()
}

func Set(key string, value string, expiration time.Duration) {
	client.Set(ctx, key, value, expiration)
}

func Get(key string) string {
	val, _ := client.Get(ctx, key).Result()

	return val
}

func Hset(key string, fieldValues map[string]string) {
	client.HSet(ctx, key, fieldValues)
}

func Expire(key string, expiration time.Duration) {
	client.Expire(ctx, key, expiration)
}

func Incr(key string) int64 {
	val, _ := client.Incr(ctx, key).Result()

	return val
}

func Del(key string) {
	client.Del(ctx, key)
}
