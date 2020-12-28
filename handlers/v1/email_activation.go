package handlers_v1

import (
	"auth/lib/config"
	"auth/lib/helpers"
	"auth/models"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
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

func GenerateActivationToken(email string, db *gorm.DB) (string, error) {
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
		Audience:  email,
		ExpiresAt: now.Add(time.Duration(config.ActivateAuth.Expire) * time.Second).Unix(),
		Issuer:    "System",
		IssuedAt:  now.Unix(),
		NotBefore: now.Add(3 * time.Second).Unix(),
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
				//todo: not evena token
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// todo : not in duration
			}
		}
		// todo unknown error
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if err := claims.Valid(); !ok || err != nil {
		// todo unknown error
	}

	//id, err := helpers.ShortStringToUuid()

	//emailVerify :=

	return nil
}
