package database

import (
	"context"
	"fmt"
	"os"
	"time"

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

func verifyRedisConnection(ctx context.Context, timeout time.Duration, rdb *redis.Client) (err error) {
	logger.Info("Init verifyRedisConnection function", zap.String("journey", "Configuration"))

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err = rdb.Ping(ctx).Err(); err != nil {
		logger.Error("Ping function returned an error", err, zap.String("journey", "Configuration"))
		return err
	}

	logger.Info("verifyRedisConnection executed successfully", zap.String("journey", "Configuration"))

	return nil
}

func RetryRedisConnection(ctx context.Context, timeout time.Duration, rdb *redis.Client, retries int) (err error) {
	logger.Info("Init RetryRedisConnection function", zap.String("journey", "Configuration"))

	for i := range retries {
		if err = verifyRedisConnection(ctx, timeout, rdb); err == nil {
			logger.Info("RetryRedisConnection executed successfully", zap.String("journey", "Configuration"))
			return nil
		} else {
			message := fmt.Sprintf("Attempt %d/%d failed", i, retries)
			logger.Error(message, err, zap.String("journey", "Configuration"))
		}

		select {
		case <-time.After(timeout):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	err = fmt.Errorf("Faield to connect with Redis database after %d attempts", retries)

	logger.Error("RetryRedisConnection finished with error", err, zap.String("journey", "Configuration"))

	return err
}
