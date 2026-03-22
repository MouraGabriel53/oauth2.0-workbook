package auth

import (
	"os"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	CLIENT_ID     = "CLIENT_ID"
	CLIENT_SECRET = "CLIENT_SECRET"
)

func NewOauth2Handler() *oauth2.Config {
	logger.Info("Init ConfigureOauth2 configuration", zap.String("journey", "Configuration"))

	return &oauth2.Config{
		ClientID:     os.Getenv(CLIENT_ID),
		ClientSecret: os.Getenv(CLIENT_SECRET),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8000/auth/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}
