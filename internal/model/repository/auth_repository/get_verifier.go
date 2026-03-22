package authrepository

import (
	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func (ar *authenticationRepositoryInterface) GetVerifier(ctx *gin.Context, state string) (stringCmd *redis.StringCmd) {
	logger.Info("Init GetVerifier function", zap.String("journey", "AuthenticateUser"))

	return ar.redis.Get(ctx, state)
}
