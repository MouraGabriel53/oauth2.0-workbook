package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var VerifierMap sync.Map

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

func generateRandomString(number int) string {
	b := make([]byte, number)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/auth/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}

	v1 := r.Group("/auth")
	{
		v1.GET("/users", func(ctx *gin.Context) {
			verifier := oauth2.GenerateVerifier()

			state := generateRandomString(32) //Generate random state 

			VerifierMap.Store(state, verifier) //Utilize REDIS

			url := conf.AuthCodeURL(state, oauth2.AccessTypeOnline, oauth2.S256ChallengeOption(verifier)) //ADD S256ChallengeOption to protect against PKCE
			ctx.Redirect(http.StatusTemporaryRedirect, url)
		})

		v1.GET("/callback", func(ctx *gin.Context) {
			code := ctx.Query("code")
			state := ctx.Query("state")

			value, ok := VerifierMap.Load(state)
			if !ok {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid state",
				})
				return
			}

			verifier := value.(string)

			VerifierMap.Delete(state)

			token, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid code",
				})
				return
			}

			client := conf.Client(ctx, token)
			resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "error to request user information",
				})
				return
			}
			defer resp.Body.Close()

			var GoogleResponse GoogleUser

			if err := json.NewDecoder(resp.Body).Decode(&GoogleResponse); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "error to decode response body",
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"user":        GoogleResponse,
				"accessToken": token.AccessToken,
				"expiry":      token.Expiry,
			})
		})
	}

	r.Run(":8080")
}
