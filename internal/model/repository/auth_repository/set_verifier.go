package authrepository

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	EXPIRATION = 0
)

func (ar *authenticationRepositoryInterface) SetVerifier(ctx *gin.Context, state, verifier string) (statusCmd *redis.StatusCmd) {
	return ar.redis.Set(ctx, state, verifier, time.Duration(EXPIRATION))
}
