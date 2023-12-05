package bootstrap

import (
	"context"
	"fmt"
	"log"

	"github.com/DitoAdriel99/go-monsterdex/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg config.Cfg) *redis.Client {

	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       0,
	})

	// Ping the Redis server to test the connection
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %s", err)
	}

	fmt.Println("Connected to Redis:", pong)
	return client
}
