package util

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

// ComparePwd 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}

func IsValidPassword(password string) bool {
	// 检查长度是否至少8位
	if len(password) < 8 {
		return false
	}

	// 检查是否包含至少一个字母
	hasLetter, _ := regexp.MatchString(`[a-zA-Z]`, password)
	if !hasLetter {
		return false
	}

	// 检查是否包含至少一个数字
	hasNumber, _ := regexp.MatchString(`[0-9]`, password)
	if !hasNumber {
		return false
	}

	return true
}
