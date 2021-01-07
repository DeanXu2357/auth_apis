package services

import (
	"auth/internal/config"
	"auth/internal/helpers"
	"auth/internal/models"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	authSignKey   *rsa.PrivateKey
	authVerifyKey *rsa.PublicKey

	ErrorUserNotFound      = errors.New("user not found")
	ErrorPasswordIncorrect = errors.New("password is incorrect")
)

func init() {
	var err error
	authSignKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(config.LoginAuth.PrivateKey))
	if err != nil {
		log.Fatal(err)
	}

	authVerifyKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(config.LoginAuth.PublicKey))
	if err != nil {
		log.Fatal(err)
	}
}

func EmailVerify(email, pwd string, db *gorm.DB) (string, error) {
	var loginInfo models.EmailLogin
	var user models.User
	result := db.Where("email = ?", email).First(&loginInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", ErrorUserNotFound
	}
	result = db.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", ErrorUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(loginInfo.Password), []byte(pwd)); err != nil {
		return "", ErrorPasswordIncorrect
	}

	a := &models.AuthToken{
		UserID:   user.ID,
		LoginWay: models.LoginByEmail,
		Revoked:  models.RevokedFalse,
	}
	if err := db.Create(a).Error; err != nil {
		return "", fmt.Errorf("Insert failed :  %w", err)
	}

	now := helpers.NowTime()
	claims := &jwt.StandardClaims{
		Audience:  loginInfo.Email,
		ExpiresAt: now.Add(time.Duration(config.LoginAuth.Expire) * time.Second).Unix(),
		Issuer:    "System",
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
		Subject:   "Login token",
		Id:        helpers.UuidToShortString(a.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(authSignKey)
}
