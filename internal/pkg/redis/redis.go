package redis

import (
	"context"
	"time"

	"github.com/redis/rueidis"
)

var client rueidis.Client
var err error
var ctx = context.Background()

func NewRedisClient(configuration string) {
	client, err = rueidis.NewClient(rueidis.MustParseURL(configuration))
	if err != nil {
		panic(err)
	}
}

func CloseRedisClient() {
	client.Close()
}

func Set(key string, value string, duration time.Duration) {
	err = client.Do(ctx, client.B().Set().Key(key).Value(value).Ex(duration).Build()).Error()
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	name, err := client.Do(ctx, client.B().Get().Key(key).Build()).ToString()
	if err != nil {
		panic(err)
	}

	return name
}

func Hset(key string, fieldValues map[string]string) {
	fieldValue := client.B().Hset().Key(key).FieldValue()
	for field, value := range fieldValues {
		fieldValue.FieldValue(field, value)
	}

	client.Do(ctx, fieldValue.Build())
}

func Expire(key string, duration time.Duration) {
	client.Do(ctx, client.B().Expire().Key(key).Seconds(int64(duration.Seconds())).Build())
}