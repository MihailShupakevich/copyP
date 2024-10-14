package db

import (
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

func RedisInit(addr, password string, db int) *redis.Client {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return RedisClient
}
