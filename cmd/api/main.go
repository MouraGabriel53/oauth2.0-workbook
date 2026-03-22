package main

import (
	"os"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration"
	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/database"
	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	authcontroller "github.com/MouraGabriel53/teste-oauth-go/internal/controller/auth_controller"
	authrepository "github.com/MouraGabriel53/teste-oauth-go/internal/model/repository/auth_repository"
	authservice "github.com/MouraGabriel53/teste-oauth-go/internal/model/service/auth_service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//MELHORAR AS CONFIGURAÇÕES (ENV) | ADICIONAR SERVER PORT, URI PARA PASSAR NO RUNTIME DO DOCKER
//ABSTRAIR INICIALIZAÇÃO DE DEPENDÊNCIAS
//ABSTRAIR CRIAÇÃO DE ROTA
//ADICIONAR LOGIN/LOGOUT COM JWT (REFRESH TOKENS)
//ADICIONAR POSTGRE NO COMPOSE
//CRIAR ROTA E CRUD DE ALGO
//CONFIGURAR DB
//OBSERVALIDADE
//VIZUALIZAÇÃO DE LOG

//CONFIGURAR ERROS OK
//CONFIGURAR LOG (UBER-ZAP) OK
//ADICIONAR REDIS NO COMPOSE OK

var (
	API_PORT = "API_PORT"
)

func main() {
	logger.Info("Initializing application")

	_ = godotenv.Load(".env")

	r := gin.Default()

	rdb := database.NewRedisClient()

	rdb.Ping()

	authrepository := authrepository.NewAuthenticationRepositoryInterface(rdb)
	authservice := authservice.NewAuthenticationServiceInterface(authrepository, googleAuth)
	authController := authcontroller.NewAuthenticationContollerInterface(authservice)

	v1 := r.Group("/auth")
	{
		v1.GET("/profile", authController.AuthenticateUser)
		v1.GET("/callback", authController.Callback)
	}

	r.Run(os.Getenv(API_PORT))
}
