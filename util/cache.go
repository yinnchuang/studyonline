package util

import "github.com/google/uuid"

func GenerateToken() string {
	token := uuid.New().String()
	return token
}
