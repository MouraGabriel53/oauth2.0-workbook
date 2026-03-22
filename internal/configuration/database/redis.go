package database

import (
	"context"
	"os"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
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

func NewRedisClient() *redis.Client {
	logger.Info("Init NewRedisClient function", zap.String("journey", "Configuration"))

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv(REDIS_ADDRESS),
		Username: os.Getenv(REDIS_USERNAME),
		Password: os.Getenv(REDIS_PASSWORD),
		DB:       REDIS_DB,
		Protocol: REDIS_PROTOCOL,
	})
}

func VerifyRedisConnection(rdb *redis.Client) (err error) {
	logger.Info("Init VerifyRedisConnection function", zap.String("journey", "Configuration"))

	ctx := context.Background()

	if err = rdb.Ping(ctx).Err(); err != nil {
		logger.Error("Ping function returned an error", err, zap.String("journey", "Configuration"))
		return err
	}

	logger.Info("VerifyRedisConnection executed successfully", zap.String("journey", "Configuration"))

	return nil
}
