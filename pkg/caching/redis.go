package redis

import (
	"context"
	"go-booking/internal/config"
	"log"

	"github.com/redis/go-redis/v9"
)

func MustConnectRedis(ctx context.Context, redisConfig config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("could not connect to redis: %v", err)
	}
	log.Println("Connected to Redis:", pong)

	return client
}
