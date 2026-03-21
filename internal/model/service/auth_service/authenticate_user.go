package authservice

import (
	"net/http"

	"github.com/MouraGabriel53/teste-oauth-go/internal/utils"
	"golang.org/x/oauth2"
)

var (
	RANDOM_STRING_LENGHT = 32
)

func (as *authenticationServiceInterface) AuthenticateUser() {
	verifier := oauth2.GenerateVerifier()

	state := utils.GenerateRandomString(RANDOM_STRING_LENGHT)

	as.repository.StoreVerifier()

	VerifierMap.Store(state, verifier) //Utilize REDIS

	url := conf.AuthCodeURL(state, oauth2.AccessTypeOnline, oauth2.S256ChallengeOption(verifier)) //ADD S256ChallengeOption to protect against PKCE
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
