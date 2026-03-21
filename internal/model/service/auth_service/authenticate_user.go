package authservice

import (
	"log/slog"
	"net/http"

	"github.com/MouraGabriel53/teste-oauth-go/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var (
	RANDOM_STRING_LENGHT = 32
)

func (as *authenticationServiceInterface) AuthenticateUser(ctx *gin.Context) {
	verifier := oauth2.GenerateVerifier()

	state := utils.GenerateRandomString(RANDOM_STRING_LENGHT)

	if statusCmd := as.repository.SetVerifier(ctx, state, verifier); statusCmd.Err() != nil {
		slog.Error("error to set verifier", "err", statusCmd.Err())
	}

	url := as.googleAuth.AuthCodeURL(state, oauth2.AccessTypeOnline, oauth2.S256ChallengeOption(verifier))
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
