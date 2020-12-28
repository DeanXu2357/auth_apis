package handlers_v1

import (
	"auth/lib/config"
	"github.com/dgrijalva/jwt-go"
)

func GenerateActivationCode(config config.LoginAuthSettings, email string) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return
}
