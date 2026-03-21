package authcontroller

import (
	"github.com/gin-gonic/gin"
)

func (ac *authenticationControllerInterface) AuthenticateUser(ctx *gin.Context) {
	ac.service.AuthenticateUser(ctx)
}
