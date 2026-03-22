package utils

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/MouraGabriel53/teste-oauth-go/internal/configuration/logger"
	"go.uber.org/zap"
)

func GenerateRandomString(number int) string {
	logger.Info("Init GenerateRandomString function", zap.String("journey", "Utils"))

	b := make([]byte, number)
	rand.Read(b)

	logger.Info("GenerateRandomString function executed successfully", zap.String("journey", "Utils"))

	return base64.URLEncoding.EncodeToString(b)
}
