package authcontroller

import (
	"net/http"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (ac *authenticationControllerInterface) Callback(ctx *gin.Context) {
	logger.Info("Init AuthenticateUser controller", zap.String("journey", "AuthenticateUser"))

	user, err := ac.service.Callback(ctx)
	if err != nil {
		logger.Error("Error trying to call Callback service", err, zap.String("journey", "AuthenticateUser"))

		ctx.JSON(err.Status, gin.H{
			"error": err,
		})
		return
	}

	logger.Info("AuthenticateUser controller executed successfully", zap.String("journey", "AuthenticateUser"))

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
