package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

func usernameID(username string) string {
	hash := sha256.Sum256([]byte(username))
	return hex.EncodeToString(hash[:])
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!,."

func generatePassword(length int) (string, error) {
	password := make([]byte, length)

	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}

	return string(password), nil
}
