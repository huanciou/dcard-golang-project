package models

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx context.Context

func RedisInit() {
	Ctx = context.Background()
	Client = NewClient(Ctx)
}

func NewClient(ctx context.Context) *redis.Client {
	ADDR := os.Getenv("REDIS_ADDR")
	PASSWORD := os.Getenv("REDIS_PASSWORD")
	DB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	client := redis.NewClient(&redis.Options{
		Addr:     ADDR,
		Password: PASSWORD,
		DB:       DB,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(pong)
	return client
}
