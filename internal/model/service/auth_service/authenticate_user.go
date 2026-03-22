package authservice

import (
	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	resterror "github.com/MouraGabriel53/teste-oauth-go/internal/configuration/rest_error"
	"github.com/MouraGabriel53/teste-oauth-go/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

var (
	RANDOM_STRING_LENGHT = 32
)

func (as *authenticationServiceInterface) AuthenticateUser(ctx *gin.Context) (url string, restError *resterror.RestError) {
	logger.Info("Init AuthenticateUser service", zap.String("journey", "AuthenticateUser"))

	verifier := oauth2.GenerateVerifier()

	state := utils.GenerateRandomString(RANDOM_STRING_LENGHT)

	if statusCmd := as.repository.SetVerifier(ctx, state, verifier); statusCmd.Err() != nil {
		logger.Error("Error trying to call SetVerifier repository", statusCmd.Err(), zap.String("journey", "AuthenticateUser"))
		return "", resterror.NewInternalServerError("couldn't set verifier")
	}

	url = as.googleAuth.AuthCodeURL(state, oauth2.AccessTypeOnline, oauth2.S256ChallengeOption(verifier))

	logger.Info("AuthenticateUser service executed successfully", zap.String("journey", "AuthenticateUser"))

	return url, nil
}
