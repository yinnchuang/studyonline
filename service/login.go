package service

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

func Login(username string, password string) (bool, string, error) {
	if username == "admin" && password == "admin" {
		token := uuid.New().String()
		token = strings.ReplaceAll(token, "-", "")
		return true, "token", nil
	} else {
		return false, "", errors.New("登录失败")
	}
}

func AdminLogin(username string, password string) (bool, string, error) {
	if username == "admin" && password == "admin" {
		
		return true, "token", nil
	} else {
		return false, "", errors.New("登录失败")
	}
}
