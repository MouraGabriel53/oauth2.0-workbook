package authservice

import (
	authrepository "github.com/MouraGabriel53/teste-oauth-go/internal/model/repository/auth_repository"
	"github.com/gin-gonic/gin"
)

func NewAuthenticationServiceInterface(repository authrepository.AuthenticationRepositoryInterface) *authenticationServiceInterface {
	return &authenticationServiceInterface{
		repository: repository,
	}
}

type AuthenticationServiceInterface interface {
	AuthenticateUser(ctx *gin.Context)
}

type authenticationServiceInterface struct {
	repository authrepository.AuthenticationRepositoryInterface
}
