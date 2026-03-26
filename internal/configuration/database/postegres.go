package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

var (
	USERNAME = "POSTGRES_USER"
	PASSWORD = "POSTGRES_PASSWORD"
	IP       = "POSTGRES_IP"
	PORT     = "POSTGRES_PORT"
	DB_NAME  = "POSTGRES_DB"
)

func NewPostgresClient() (db *pgx.Conn, err error) {
	ctx := context.Background()

	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv(USERNAME), os.Getenv(PASSWORD), os.Getenv(IP), os.Getenv(PORT), os.Getenv(DB_NAME))

	db, err = pgx.Connect(ctx, url)

	return db, err
}

func verifyPostegresConnection(ctx context.Context, db *pgx.Conn, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return db.Ping(ctx)
}

func RetryPostgresConnection(ctx context.Context, timeout time.Duration, db *pgx.Conn, retries int) (err error) {
	logger.Info("Init RetryRedisConnection function", zap.String("journey", "Configuration"))

	for i := range retries {
		if err = verifyPostegresConnection(ctx, db, timeout); err == nil {
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
