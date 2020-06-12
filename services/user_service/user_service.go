package user_service

import (
	"auth/models"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func New() (service *UserService,err error) {
	return
}

func (service *UserService)GetUserByUUID(uuid string) (err error) {
	return
}

func (service *UserService)Create(data map[string]interface{}) (err error) {
	user := &models.User{
		Name: data["name"].(string),
		Email: data["email"].(string),
	}
	user.SetConnection(service.DB)
	if err = user.Create(); err != nil {
		return
	}

	return
}
