package user_service

import (
	"auth/app"
	"auth/models"
)

type UserService struct {
	app *app.Instance
}

func New(app *app.Instance) *UserService {
	return &UserService{app: app}
}

func (s *UserService)GetUserByUUID(uuid string) (user ,err error) {
	return
}

func (s *UserService)Create(data map[string]interface{}) (err error) {
	user := &models.User{
		Name: data["name"].(string),
		Email: data["email"].(string),
	}
	user.SetConnection(s.app.Database)
	if err = user.Create(); err != nil {
		return
	}

	return
}
