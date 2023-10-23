package initDB

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var REDIS_DB *redis.Client

func InitRedis() {
	REDIS_DB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	pong, err := REDIS_DB.Ping(context.Background()).Result()
	fmt.Println(pong, err)
	if err != nil {
		panic(err)
	}
}
