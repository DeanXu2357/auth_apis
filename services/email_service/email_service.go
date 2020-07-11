package email_service

import (
	"auth/app"
	"auth/models"
)

type EmailService struct {
	app *app.Instance
}

func New(app *app.Instance) *EmailService {
	return &EmailService{app: app}
}

func (s *EmailService)RegistByMail(email string, name string, password string) (err error) {
	// transaction create user create email_certificates

	// send email

	return
}

func (s *EmailService)ResendMail(email string) (err error) {
	// send email

	return
}

func (s *EmailService)sendMail(email string, text string) (err error) {
	return
}

func ActivateEmailRegister(email string, token string) (err error) {
	return
}

func EmailLogin(email string, pwd string) (user models.User, err error) {
	return
}
