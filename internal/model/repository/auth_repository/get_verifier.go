package authrepository

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func (ar *authenticationRepositoryInterface) GetVerifier(ctx *gin.Context, state string) (stringCmd *redis.StringCmd) {
	return ar.redis.Get(ctx, state)
}
