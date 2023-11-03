package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

var RedisImpl Cache

type Cache struct {
	client *redis.Client
}

func Init() {
	fmt.Println("Redis local mode")
	client := redis.NewClient(&redis.Options{
		Addr:     "cache",
		Password: "",
		DB:       0,
	})
	if _, err := client.Ping().Result(); err != nil {
		errMsg := fmt.Sprintf("Failed to connect to cache: %s", err)
		panic(errMsg)
	}
	RedisClient = client
	RedisImpl = Cache{client: RedisClient}
}
