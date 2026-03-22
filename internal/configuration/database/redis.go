package database

import (
	"context"
	"os"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	resterror "github.com/MouraGabriel53/teste-oauth-go/internal/configuration/rest_error"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	REDIS_ADDRESS  = "REDIS_ADDRESS"
	REDIS_USERNAME = "REDIS_USERNAME"
	REDIS_PASSWORD = "REDIS_PASSWORD"
	REDIS_DB       = 0
	REDIS_PROTOCOL = 2
)

type RedisClient struct {
	rdb *redis.Client
}

func NewRedisClient() *RedisClient {
	logger.Info("Init NewRedisClient database", zap.String("journey", "Configuration"))

	return &RedisClient{
		rdb: redis.NewClient(&redis.Options{
			Addr:     os.Getenv(REDIS_ADDRESS),
			Username: os.Getenv(REDIS_USERNAME),
			Password: os.Getenv(REDIS_PASSWORD),
			DB:       REDIS_DB,
			Protocol: REDIS_PROTOCOL,
		}),
	}
}

func (rc *RedisClient) Ping() {
	logger.Info("Init Ping redis database", zap.String("journey", "Configuration"))

	ctx := context.Background()

	if statusCmd := rc.rdb.Ping(ctx); statusCmd.Err() != nil {
		logger.Error("Error trying to call Ping", statusCmd.Err(), zap.String("journey", "Configuration"))
		resterror.NewInternalServerError("Error trying to connect with Redis")
	}

	logger.Info("Ping redis databse executed successfully", zap.String("journey", "Configuration"))
}
