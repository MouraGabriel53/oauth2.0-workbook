package authservice

import (
	authrepository "github.com/MouraGabriel53/teste-oauth-go/internal/model/repository/auth_repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func NewAuthenticationServiceInterface(repository authrepository.AuthenticationRepositoryInterface, googleAuth *oauth2.Config) *authenticationServiceInterface {
	return &authenticationServiceInterface{
		repository: repository,
		googleAuth: googleAuth,
	}
}

type AuthenticationServiceInterface interface {
	AuthenticateUser(ctx *gin.Context)
}

type authenticationServiceInterface struct {
	repository authrepository.AuthenticationRepositoryInterface
	googleAuth *oauth2.Config
}
