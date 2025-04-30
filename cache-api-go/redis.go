package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx = context.Background()
var rdb *redis.Client

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",           // default Redis port
		Password: "your_redis_password_here", // no password set
		DB:       0,                          // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
