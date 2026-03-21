package configuration

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	REDIS_ADDRESS  = "REDIS_ADDRESS"
	REDIS_PASSWORD = "REDIS_PASSWORD"
	REDIS_DB       = 0
	REDIS_PROTOCOL = 2
)

func ConfigureRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv(REDIS_ADDRESS),
		Password: os.Getenv(REDIS_PASSWORD),
		DB:       REDIS_DB,
		Protocol: REDIS_PROTOCOL,
	})
}
