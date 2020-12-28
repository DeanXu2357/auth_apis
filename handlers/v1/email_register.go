package handlers_v1

import (
	"auth/models"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
)

func Register(name string, email string, password string, db *gorm.DB) (*models.User, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Print(r.(error))
			tx.Rollback()
		}
	}()

	user := &models.User{Name: name, Email: email}
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, ErrEmailAlreadyRegistered
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error due to :  %w", err)
	}

	if err := tx.Create(&models.EmailLogin{Email: email, Password: string(hashedPassword)}).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error due to :  %w", err)
	}

	tx.Commit()

	return user, nil
}
