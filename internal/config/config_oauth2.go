package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	CLIENT_ID     = "CLIENT_ID"
	CLIENT_SECRET = "CLIENT_SECRET"
)

func ConfigurateOauth2() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv(CLIENT_ID),
		ClientSecret: os.Getenv(CLIENT_SECRET),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/auth/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}
