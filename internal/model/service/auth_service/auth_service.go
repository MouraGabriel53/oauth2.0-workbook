package authservice

import (
	resterror "github.com/MouraGabriel53/teste-oauth-go/internal/configuration/rest_error"
	authrepository "github.com/MouraGabriel53/teste-oauth-go/internal/model/repository/auth_repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthenticationServiceInterface interface {
	AuthenticateUser(ctx *gin.Context) (url string, restError *resterror.RestError)
	Callback(ctx *gin.Context) (GoogleResponse *GoogleUser, restError *resterror.RestError)
}

type authenticationServiceInterface struct {
	repository authrepository.AuthenticationRepositoryInterface
	googleAuth *oauth2.Config
}

func NewAuthenticationServiceInterface(repository authrepository.AuthenticationRepositoryInterface, googleAuth *oauth2.Config) *authenticationServiceInterface {
	return &authenticationServiceInterface{
		repository: repository,
		googleAuth: googleAuth,
	}
}
