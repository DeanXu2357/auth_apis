package handlers_v1

import (
	"auth/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

func Register(name string, email string, password string, db *gorm.DB) (*models.User, error) {
	// todo : find a transaction manager library
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
		return nil, errors.New(EmailAlreadyRegistered)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return nil, err
	}

	if err := tx.Create(&models.EmailLogin{Email: email, Password: string(hashedPassword)}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return user, nil
}

// lpush key_name value
func dispatchMailQueue(email string, mType string) error {
	return nil
}
