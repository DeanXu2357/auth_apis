package handlers_v1

import (
	"auth/internal/models"
	"gorm.io/gorm"
)

func FindEmailLogin(email string, db *gorm.DB) (models.EmailLogin, error) {
	// todo : not verify error

	return models.EmailLogin{}, nil
}
