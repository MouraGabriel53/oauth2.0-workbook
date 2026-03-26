package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/auth"
	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/database"
	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	authcontroller "github.com/MouraGabriel53/teste-oauth-go/internal/controller/auth_controller"
	"github.com/MouraGabriel53/teste-oauth-go/internal/middleware"
	authrepository "github.com/MouraGabriel53/teste-oauth-go/internal/model/repository/auth_repository"
	authservice "github.com/MouraGabriel53/teste-oauth-go/internal/model/service/auth_service"
	"github.com/gin-gonic/gin"
)

//ABSTRAIR INICIALIZAÇÃO DE DEPENDÊNCIAS
//ABSTRAIR CRIAÇÃO DE ROTA
//ADICIONAR LOGIN/LOGOUT COM JWT (REFRESH TOKENS)
//ADICIONAR POSTGRE NO COMPOSE
//CRIAR ROTA E CRUD DE ALGO
//CONFIGURAR DB
//OBSERVALIDADE
//VIZUALIZAÇÃO DE LOG
//CONFIGURAR TIMEOUT, RATE LIMITE E GRACEFULL SHUTDOWN
//CONFIGURAR LIMITES DE CONEXOES COM GIN, REDIS, POSTGRES

//MELHORAR AS CONFIGURAÇÕES (ENV) | ADICIONAR SERVER PORT, URI PARA PASSAR NO RUNTIME DO DOCKER OK
//CONFIGUURAR HEALTHCHEKC, TRY PARA SUBIR O REDIS OK
//CONFIGURAR ERROS OK
//CONFIGURAR LOG (UBER-ZAP) OK
//ADICIONAR REDIS NO COMPOSE OK

var (
	// ENV_PATH      = "../../../.env"
	GIN_MODE      = "GIN_MODE"
	API_PORT      = "DOCKER_API_PORT"
	ALLOW_ORIGINS = []string{"*"}
)

func main() {
	logger.Info("Initializing application")

	// _ = godotenv.Load(ENV_PATH)

	//===============================================================================

	gin.SetMode(os.Getenv(GIN_MODE))

	r := gin.Default()

	r.Use(middleware.NewCorsHandler(ALLOW_ORIGINS))

	//===============================================================================

	oauth2Handler := auth.NewOauth2Handler()

	//===============================================================================

	rdb := database.NewRedisClient()

	ctx := context.Background()

	timeout := time.Duration(10 * time.Second)

	retries := 5

	if err := database.RetryRedisConnection(ctx, timeout, rdb, retries); err != nil {
		panic(err)
	}

	//===============================================================================

	db, err := database.NewPostgresClient()
	if err != nil {
		panic(err)
	}
	defer db.Close(ctx)

	if err := database.RetryPostgresConnection(ctx, timeout, db, retries); err != nil {
		panic(err)
	}

	//===============================================================================

	authrepository := authrepository.NewAuthenticationRepositoryInterface(rdb)
	authservice := authservice.NewAuthenticationServiceInterface(authrepository, oauth2Handler)
	authController := authcontroller.NewAuthenticationContollerInterface(authservice)

	//===============================================================================

	v1 := r.Group("/auth")
	{
		v1.GET("/profile", authController.AuthenticateUser)
		v1.GET("/callback", authController.Callback)
	}

	//===============================================================================

	apiPort := fmt.Sprintf(":%s", os.Getenv(API_PORT))

	r.Run(apiPort)

	logger.Info("Exiting gracefully")
}
