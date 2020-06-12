package email_service

import (
    "auth/services/user_service"
	"github.com/jinzhu/gorm"
)

type EmailService struct {
    DB *gorm.DB
}

func RegistByMail(email string) (err error) {
	// transaction create user create email_certificates

	// send email

	return
}

func ResendMail(email string) (err error) {
	// send email

	return
}

func sendMail(email string, text string) (err error) {
	return
}

func ActivateEmailRegister(email string, token string) (err error) {
	return
}

func EmailLogin(email string, pwd string) (user user_service.User, err error) {
	return
}
