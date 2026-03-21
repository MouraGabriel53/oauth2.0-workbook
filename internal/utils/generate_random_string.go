package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(number int) string {
	b := make([]byte, number)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
