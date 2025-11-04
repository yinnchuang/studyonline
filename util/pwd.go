package util

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
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

func IsValidEmail(email string) bool {
	// 1. 长度检查（按常见邮箱长度 6~254 做简单过滤）
	if len(email) < 6 || len(email) > 254 {
		return false
	}

	// 2. 正则校验
	//    ^[a-zA-Z0-9._-]+        用户名
	//    @[a-zA-Z0-9-]+(\.[a-zA-Z]{2,})+$  域名
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9-]+(\.[a-zA-Z]{2,})+$`, email)
	return ok
}
