package handlers_v1

import (
	"auth/internal/models"
	"gorm.io/gorm"
)

func GeneratePasswordToken(user *models.User, db *gorm.DB) (string, error) {

}
