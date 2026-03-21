package authrepository

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func NewAuthenticationRepositoryInterface(redis *redis.Client) *authenticationRepositoryInterface {
	return &authenticationRepositoryInterface{
		redis: redis,
	}
}

type AuthenticationRepositoryInterface interface {
	SetVerifier(ctx *gin.Context, state, verifier string) (statusCmd *redis.StatusCmd)
	GetVerifier()
}

type authenticationRepositoryInterface struct {
	redis *redis.Client
}
