package handlers_v1

import (
	"auth/lib/config"
	"auth/lib/helpers"
	"auth/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"time"
)

func GenerateActivationCode(config config.ActivateAuthSettings, email string, db *gorm.DB) (string, error) {
	v := &models.EmailVerify{
		Email:        email,
		MailType:     models.MailTypeVerifyAccount,
		Verification: models.VerifyFalse,
	}
	if err := db.Create(v).Error; err != nil {
		return "", fmt.Errorf("Insert failed :  %w", err)
	}

	now := time.Now()
	claims := &jwt.StandardClaims{
		Audience: email,
		ExpiresAt: now.Add(time.Duration(config.Expire) * time.Second).Unix(),
		Issuer:    "System",
		IssuedAt:  now.Unix(),
		NotBefore: now.Add(3 * time.Second).Unix(),
		Subject:   "Account email verification",
		Id: helpers.UuidToShortString(v.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString([]byte(config.Secret))
}
