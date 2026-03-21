package authcontroller

import (
	authservice "github.com/MouraGabriel53/teste-oauth-go/internal/model/service/auth_service"
	"github.com/gin-gonic/gin"
)

func NewAuthenticationContollerInterface(service authservice.AuthenticationServiceInterface) *authenticationControllerInterface {
	return &authenticationControllerInterface{
		service: service,
	}
}

type AuthenticationControllerInterface interface {
	AuthenticateUser(ctx *gin.Context)
	Callback(ctx *gin.Context)
}

type authenticationControllerInterface struct {
	service authservice.AuthenticationServiceInterface
}
