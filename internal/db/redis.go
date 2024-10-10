package db

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	ctx         = context.Background()
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

func RedisAddData() {
	err := RedisClient.HSet(ctx, "user_profile", map[string]interface{}{
		"name":  "John Doe",
		"age":   30,
		"email": "john.doe@example.com",
	}).Err()
	if err != nil {
		panic(err)
	}
}
