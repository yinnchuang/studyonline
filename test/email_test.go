package test

import (
	"fmt"
	"net/smtp"
	"testing"
)

func TestSend(t *testing.T) {
	// 邮件服务器信息
	smtpServer := "smtp.imicams.ac.cn" // SMTP 服务器地址
	smtpPort := "587"                  // SMTP 服务器端口号
	smtpUser := "zsgc@imicams.ac.cn"   // 发件人邮箱地址
	smtpPassword := "Znyx#25117"       // 发件人邮箱密码

	// 邮件内容
	from := smtpUser
	to := []string{"yinnchuang@163.com"} // 收件人邮箱地址
	subject := "Test Email from Go"
	body := "This is a test email sent from Go using SMTP."

	// 构造邮件内容
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to[0], subject, body)

	// 连接到 SMTP 服务器
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		fmt.Printf("Error sending email: %s\n", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
