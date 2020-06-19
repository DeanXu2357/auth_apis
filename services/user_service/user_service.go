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

func (service *UserService)GetUserByUUID(uuid string) (err error) {
	return
}

func (service *UserService)Create(data map[string]interface{}) (err error) {
	user := &models.User{
		Name: data["name"].(string),
		Email: data["email"].(string),
	}
	user.SetConnection(service.app.Database)
	if err = user.Create(); err != nil {
		return
	}

	return
}
