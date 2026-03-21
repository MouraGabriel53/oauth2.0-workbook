package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration"
	authcontroller "github.com/MouraGabriel53/teste-oauth-go/internal/controller/auth_controller"
	"github.com/MouraGabriel53/teste-oauth-go/internal/controller/routes"
	authrepository "github.com/MouraGabriel53/teste-oauth-go/internal/model/repository/auth_repository"
	authservice "github.com/MouraGabriel53/teste-oauth-go/internal/model/service/auth_service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
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

func main() {

	_ = godotenv.Load(".env")

	r := gin.Default()

	googleAuth := configuration.ConfigureOauth2()

	redis := configuration.ConfigureRedisClient()

	authrepository := authrepository.NewAuthenticationRepositoryInterface(redis)
	authservice := authservice.NewAuthenticationServiceInterface(authrepository, googleAuth)
	authController := authcontroller.NewAuthenticationContollerInterface(authservice)

	routes.AuthRoutes(r, authController.AuthenticateUser)

	v1 := r.Group("/auth")
	{
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

			token, err := googleAuth.Exchange(ctx, code, oauth2.VerifierOption(verifier))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid code",
				})
				return
			}

			client := googleAuth.Client(ctx, token)
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

	r.Run(":8000")
}
