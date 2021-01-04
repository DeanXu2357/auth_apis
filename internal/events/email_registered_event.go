package events

import (
	"auth/internal/models"
	"fmt"
)

type EmailRegisteredEvent struct {
	User  models.User
	Token string
}

func (e EmailRegisteredEvent) GetName() string {
	return EmailRegistered
}

func (e EmailRegisteredEvent) GetUser() models.User {
	return e.User
}

func (e EmailRegisteredEvent) To() string {
	return e.User.Email
}

func (e EmailRegisteredEvent) GetSubject() string {
	return "Please identify your email address"
}

func (e EmailRegisteredEvent) GetBody() string {
	return fmt.Sprintf("click here to verify your email address %s", e.Token)
}
