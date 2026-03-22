package main

import (
	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration"
	authcontroller "github.com/MouraGabriel53/teste-oauth-go/internal/controller/auth_controller"
	authrepository "github.com/MouraGabriel53/teste-oauth-go/internal/model/repository/auth_repository"
	authservice "github.com/MouraGabriel53/teste-oauth-go/internal/model/service/auth_service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//MELHORAR AS CONFIGURAÇÕES (ENV) | ADICIONAR SERVER PORT, URI PARA PASSAR NO RUNTIME DO DOCKER
//ABSTRAIR INICIALIZAÇÃO DE DEPENDÊNCIAS
//ABSTRAIR CRIAÇÃO DE ROTAS
//ADICIONAR LOGIN/LOGOUT COM JWT (REFRESH TOKENS)
//ADICIONAR REDIS NO COMPOSE
//ADICIONAR POSTGRE NO COMPOSE
//CRIAR ROTA E CRUD DE ALGO
//CONFIGURAR LOG (UBER-ZAP)
//CONFIGURAR DB
//CONFIGURAR ERROS

func main() {

	_ = godotenv.Load(".env")

	r := gin.Default()

	googleAuth := configuration.ConfigureOauth2()

	redis := configuration.ConfigureRedisClient()

	authrepository := authrepository.NewAuthenticationRepositoryInterface(redis)
	authservice := authservice.NewAuthenticationServiceInterface(authrepository, googleAuth)
	authController := authcontroller.NewAuthenticationContollerInterface(authservice)

	v1 := r.Group("/auth")
	{
		v1.GET("/profile", authController.AuthenticateUser)
		v1.GET("/callback", authController.Callback)
	}

	r.Run(":8000")
}
