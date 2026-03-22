package authrepository

import (
	"time"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	EXPIRATION = 5 * time.Minute
)

func (ar *authenticationRepositoryInterface) SetVerifier(ctx *gin.Context, state, verifier string) (statusCmd *redis.StatusCmd) {
	logger.Info("Init SetVerifier", zap.String("journey", "AuthenticateUser"))

	return ar.redis.Set(ctx, state, verifier, time.Duration(EXPIRATION))
}
