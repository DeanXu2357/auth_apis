package email

import (
	"crypto/tls"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"log"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host string
	Port int
	IsSSl bool
	UserName string
	Password string
	From string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{info}
}

func (e *Email) SendMail(from string, to []string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("from", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSl}
	return dialer.DialAndSend(m)
}

func ReadConfig() *SMTPInfo {
	i := &SMTPInfo{}
	err := viper.Unmarshal(i)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
