package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortCode(url string) string {
	hash := sha256.Sum256([]byte(url))
	encoded := base64.RawURLEncoding.EncodeToString(hash[:])
	return encoded[:8]
}
