package authcontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (ac *authenticationControllerInterface) Callback(ctx *gin.Context) {
	user := ac.service.Callback(ctx)

	fmt.Println(user)
}
