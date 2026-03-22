package authservice

import (
	"encoding/json"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	resterror "github.com/MouraGabriel53/teste-oauth-go/internal/configuration/rest_error"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

var (
	URL = "https://www.googleapis.com/oauth2/v2/userinfo"
)

func (as *authenticationServiceInterface) Callback(ctx *gin.Context) (GoogleResponse *GoogleUser, restError *resterror.RestError) {
	logger.Info("Init Callback service", zap.String("journey", "AuthenticateUser"))

	code := ctx.Query("code")
	state := ctx.Query("state")

	redisVerifier := as.repository.GetVerifier(ctx, state)

	verifier, err := redisVerifier.Result()
	if err != nil {
		logger.Error("GetVerifier returned an error", err, zap.String("journey", "AuthenticateUser"))
		return &GoogleUser{}, resterror.NewInternalServerError("couldn't get verifier from Redis")
	}

	token, err := as.googleAuth.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		logger.Error("googleAuth.Exchange returned an error", err, zap.String("journey", "AuthenticateUser"))
		return &GoogleUser{}, resterror.NewInternalServerError("couldn't convert an authorization code into a token")
	}

	client := as.googleAuth.Client(ctx, token)
	resp, err := client.Get(URL)
	if err != nil {
		logger.Error("client.Get returned an error", err, zap.String("journey", "AuthenticateUser"))
		return &GoogleUser{}, resterror.NewInternalServerError("couldn't request user information")
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&GoogleResponse); err != nil {
		logger.Error("json.NewDecoder returned an error", err, zap.String("journey", "AuthenticateUser"))
		return &GoogleUser{}, resterror.NewInternalServerError("couldn't read response user information")
	}

	logger.Info("Callback service executed successfully", zap.String("journey", "AuthenticateUser"))

	return GoogleResponse, nil
}
