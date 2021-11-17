package db

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func CheckPasswordHash(password, hash string) bool {
	return HashPassword(password) == hash
}
