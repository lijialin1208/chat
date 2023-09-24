package initDB

import "github.com/redis/go-redis/v9"

var REDIS_DB *redis.Client

func InitRedis() {
	REDIS_DB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
}
