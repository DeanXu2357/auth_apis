package events

import "auth/models"

type EmailRegisteredEvent struct {
	user models.User
}

func (e EmailRegisteredEvent) GetName() string {
	return EmailRegistered
}

func (e EmailRegisteredEvent) GetUser() models.User {
	return e.user
}
