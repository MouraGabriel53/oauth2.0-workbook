package authservice

import (
	"encoding/json"
	"log/slog"

	"github.com/gin-gonic/gin"
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
	URL            = "https://www.googleapis.com/oauth2/v2/userinfo"
)

func (as *authenticationServiceInterface) Callback(ctx *gin.Context) (GoogleResponse *GoogleUser) {
	code := ctx.Query("code")
	state := ctx.Query("state")

	redisVerifier := as.repository.GetVerifier(ctx, state)

	verifier, err := redisVerifier.Result()
	if err != nil {
		slog.Error("error to get verifier from redis", "err", err)
	}

	token, err := as.googleAuth.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		slog.Error("invalid code", "err", err)
	}

	client := as.googleAuth.Client(ctx, token)
	resp, err := client.Get(URL)
	if err != nil {
		slog.Error("error to request user information", "err", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&GoogleResponse); err != nil {
		slog.Error("error to read body", "err", err)
	}

	return GoogleResponse
}
