package util

import (
	"fmt"
	"log"
	"net/smtp"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"studyonline/config"
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

func SendCodeToEmail(email string, code string) error {
	smtpConfig := config.AppConfig.SMTP

	// 打印配置用于调试（隐藏密码）
	log.Printf("SMTP Config - Server: %s, Port: %s, User: %s, PasswordLen: %d",
		smtpConfig.Server, smtpConfig.Port, smtpConfig.User, len(smtpConfig.Password))

	// 邮件内容
	from := smtpConfig.User
	to := []string{email} // 收件人邮箱地址
	subject := "找回密码验证码"
	body := "本次验证码为：" + code + "\n" + "验证码有效时长5分钟"

	// 构造邮件内容
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to[0], subject, body)

	// 连接到 SMTP 服务器
	auth := smtp.PlainAuth("", smtpConfig.User, smtpConfig.Password, smtpConfig.Server)
	err := smtp.SendMail(smtpConfig.Server+":"+smtpConfig.Port, auth, from, to, []byte(message))
	if err != nil {
		log.Printf("Error sending email: %s\n", err)
		return err
	}

	return nil
}
