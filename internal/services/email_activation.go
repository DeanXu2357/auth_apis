package services

import (
	"auth/internal/config"
	"auth/internal/helpers"
	"auth/internal/models"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey

	ErrorDBUpdateFailed   = errors.New("DB update failed")
	ErrorDBInsertFailed   = errors.New("DB insert failed")
	ErrorTokenMalformed   = errors.New("token malformed")
	ErrorTokenExpired     = errors.New("token expired")
	ErrorTokenNotValidYet = errors.New("token not in valid period yet")
	ErrorTokenInvalid     = errors.New("invalid token")
)

func init() {
	var err error
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(config.ActivateAuth.PrivateKey))
	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(config.ActivateAuth.PublicKey))
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateActivationToken(user models.User, db *gorm.DB) (string, error) {
	v := &models.EmailVerify{
		Email:        user.Email,
		MailType:     models.MailTypeVerifyAccount,
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
		Subject:   "Account email verification",
		Id:        helpers.UuidToShortString(v.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(signKey)
}

func Activate(tokenString string, db *gorm.DB) error {
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

	jti := claims["jti"].(string)
	id, err := helpers.ShortStringToUuid(jti)

	emailVerify := models.EmailVerify{ID: id}
	result := db.First(&emailVerify)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("data raw not exist : %w", ErrorTokenInvalid)
	}

	// todo: check validation is false
	// todo: check type is MailTypeVerifyAccount

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
	if err := tx.Model(&models.EmailLogin{Email: emailVerify.Email}).Update("verified_at", helpers.NowTime()).Error; err != nil {
		tx.Rollback()
		return ErrorDBUpdateFailed
	}
	tx.Commit()

	return nil
}
