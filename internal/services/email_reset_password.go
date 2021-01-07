package handlers_v1

import (
	"auth/internal/config"
	"auth/internal/helpers"
	"auth/internal/models"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

func GeneratePasswordToken(user models.User, db *gorm.DB) (string, error) {
	v := &models.EmailVerify{
		Email:        user.Email,
		MailType:     models.MailTypeResetPassword,
		Verification: models.VerifyFalse,
		UserID:       user.ID,
	}
	if err := db.Create(v).Error; err != nil {
		return "", fmt.Errorf("Insert failed :  %w", err)
	}

	now := helpers.NowTime()
	claims := &jwt.StandardClaims{
		Audience:  user.Email,
		ExpiresAt: now.Add(time.Duration(config.ActivateAuth.Expire) * time.Second).Unix(),
		Issuer:    "System",
		IssuedAt:  now.Unix(),
		NotBefore: now.Add(10 * time.Second).Unix(),
		Subject:   "Account password reset",
		Id:        helpers.UuidToShortString(v.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(signKey)
}

func Reset(tokenString, newPwd string, db *gorm.DB) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return ErrorTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return ErrorTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return ErrorTokenNotValidYet
			}
		}
		return ErrorTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if err := claims.Valid(); !ok || err != nil {
		return ErrorTokenInvalid
	}

	id, err := helpers.ShortStringToUuid(claims["jti"].(string))

	emailVerify := models.EmailVerify{ID: id}
	result := db.First(&emailVerify)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("data raw not exist : %w", ErrorTokenInvalid)
	}

	if emailVerify.MailType != models.MailTypeResetPassword {
		// todo : error
	}
	if emailVerify.Verification != models.VerifyFalse {
		// todo : error duplicate reset
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPwd), 8)
	if err != nil {
		return fmt.Errorf("Error due to :  %w", err)
	}

	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer func() {
		if r := recover(); r != nil {
			log.Print(r.(error))
			tx.Rollback()
		}
	}()
	if err := tx.Model(&emailVerify).Update("Verification", models.VerifyTrue).Error; err != nil {
		tx.Rollback()
		return ErrorDBUpdateFailed
	}
	if err := tx.Model(&models.EmailLogin{Email: emailVerify.Email}).Updates(models.EmailLogin{Password: string(hashedPassword)}).Error; err != nil {
		tx.Rollback()
		return ErrorDBUpdateFailed
	}
	tx.Commit()

	return nil
}
