package events

import "auth/models"

type EmailRegisteredEvent struct {
	user models.User
}

func NewEmailRegisteredEvent(u models.User) EmailRegisteredEvent {
	return EmailRegisteredEvent{user: u}
}

func (e EmailRegisteredEvent) GetName() string {
	return EmailRegistered
}

func (e EmailRegisteredEvent) GetUser() models.User {
	return e.user
}

func (e EmailRegisteredEvent) To() string {
	return e.user.Email
}

func (e EmailRegisteredEvent) GetSubject() string {
	return "Please identify your email address"
}

func (e EmailRegisteredEvent) GetBody() string {
	return "click here to verify your email address"
}
