package middleware

import (
	"time"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewCorsHandler(allowOrigins []string) gin.HandlerFunc {
	logger.Info("Init NewCorsConfig function", zap.String("journey", "Configuration"))

	return cors.New(cors.Config{
		AllowOrigins:  allowOrigins,
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	})
}
