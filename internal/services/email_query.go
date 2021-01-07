package services

import (
	"auth/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

var (
	ErrLoginNotExist = errors.New("login way not exist")
	ErrInternalError = errors.New("internal error")
	ErrEmailNotVerify = errors.New("email not verify yet")
)

func FindEmailLogin(email string, db *gorm.DB) (models.EmailLogin, error) {
	var emailLogin models.EmailLogin
	if err := db.Where(&models.EmailLogin{Email: email}).First(&emailLogin).Error; err != nil {
		if errors.Is(err , gorm.ErrRecordNotFound) {
			return emailLogin, ErrLoginNotExist
		}

		log.Print(err.Error())
		return emailLogin, fmt.Errorf("%s \n%w", err.Error(), ErrInternalError)
	}

	if emailLogin.VerifiedAt.IsZero() {
		return emailLogin, ErrEmailNotVerify
	}

	return emailLogin, nil
}
