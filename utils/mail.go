package utils

import (
	"blog/consts"
	"blog/global"
	"crypto/rand"
	"math/big"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func SendEmail(to []string, subject string, body string) error {
	mailConf := global.Conf.Email
	m := gomail.NewMessage()
	// 设置发件人
	m.SetHeader("From", mailConf.From)
	// 设置收件人
	m.SetHeader("To", to...)
	// 设置主题
	m.SetHeader("Subject", subject)
	// 设置内容
	m.SetBody("text/html", body)
	// 建立SMTP连接
	d := gomail.NewDialer(mailConf.Host, mailConf.Port, mailConf.Username, mailConf.Password)
	// 发送文件
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	logrus.Infof(consts.SendMailSuccess)
	return nil
}

// GenerateSecureCode 生成6位安全随机数字验证码
func GenerateSecureCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		code += n.String()
	}
	return code
}
