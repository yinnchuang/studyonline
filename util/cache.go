package util

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func GenerateToken() string {
	token := uuid.New().String()
	return token
}

func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(90000) + 10000
	return strconv.Itoa(code)
}
