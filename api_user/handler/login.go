package handler

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

