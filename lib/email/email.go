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
	Host     string
	Port     int
	IsSSl    bool   `mapstructure:"is_ssl"`
	UserName string `mapstructure:"user_name"`
	Password string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{info}
}

func (e *Email) SendMail(to []string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSl}
	return dialer.DialAndSend(m)
}

func NewInfo() *SMTPInfo {
	i := &SMTPInfo{}
	err := viper.UnmarshalKey("email", &i)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
