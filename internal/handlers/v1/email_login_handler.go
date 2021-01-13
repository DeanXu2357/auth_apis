package handlers_v1

import (
	"auth/internal/config"
	"auth/internal/events"
	"auth/internal/helpers"
	"auth/internal/models"
	"auth/internal/services"
	"auth/lib/email"
	"auth/lib/event_listener"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"time"
)

func RegisterByMail(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, nil)
		return
	}

	db := helpers.GetDB(c)
	user, err := services.Register(input.Name, input.Email, input.Password, db.Session(&gorm.Session{NewDB: true}))
	if err != nil {
		if errors.Is(err, services.ErrEmailAlreadyRegistered) {
			helpers.GenerateResponse(c, helpers.ReturnDuplicate, map[string]interface{}{"msg": err.Error()})
			return
		}

		helpers.GenerateResponse(c, helpers.ReturnInternalError, map[string]interface{}{"msg": err.Error()})
		return
	}

	newSession := db.Session(&gorm.Session{NewDB: true})
	token, err := services.GenerateActivationToken(*user, newSession)
	if err != nil {
		helpers.GenerateResponse(c, helpers.ReturnInternalError, err.Error())
		return
	}

	dispatcher := c.MustGet("Dispatcher").(*event_listener.Dispatcher)
	dispatcher.DispatchAsync(events.EmailRegisteredEvent{User: *user, Token: token})

	helpers.GenerateResponse(c, helpers.ReturnOK, map[string]interface{}{"user_id": user.ID, "email": input.Email})
	return
}

func VerifyMailLogin(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, nil)
		return
	}

	db := helpers.GetDB(c)
	token, err := services.EmailVerify(input.Email, input.Password, db)
	if err != nil {
		if errors.Is(err, services.ErrorUserNotFound) {
			helpers.GenerateResponse(c, helpers.ReturnNotExist, nil)
			return
		}

		helpers.GenerateResponse(c, helpers.ReturnBadRequest, map[string]string{"detail": err.Error()})
		return
	}

	helpers.GenerateResponse(c, helpers.ReturnOK, map[string]string{"token": token})
	return
}

func ActivateEmailRegister(c *gin.Context) {
	var input struct {
		Token string `json:"token"`
	}
	if err := c.ShouldBindQuery(&input); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, nil)
		return
	}

	db := helpers.GetDB(c)

	if err := services.Activate(input.Token, db); err != nil {
		if errors.Is(err, services.ErrorTokenNotValidYet) {
			helpers.GenerateResponse(c, helpers.ReturnTokenExpire, err)
			return
		} else if errors.Is(err, services.ErrorTokenExpired) {
			helpers.GenerateResponse(c, helpers.ReturnTokenExpire, err)
			return
		} else if errors.Is(err, services.ErrorTokenMalformed) {
			helpers.GenerateResponse(c, helpers.ReturnValidationFailed, err)
			return
		}

		helpers.GenerateResponse(c, helpers.ReturnInvalidToken, err)
		return
	}

	helpers.GenerateResponse(c, helpers.ReturnOK, nil)
	return
}

func RecoveryPassword(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindQuery(&input); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, nil)
		return
	}

	db := helpers.GetDB(c)

	emailLogin, err := services.FindEmailLogin(input.Email, db)
	if err != nil {
		return
	}

	var user models.User
	db.Where(&models.User{Email: emailLogin.Email}).First(&user)

	token, err := services.GeneratePasswordToken(user, db)
	if err != nil {
		// todo handle error
	}

	if err := email.SendMail(
		[]string{input.Email},
		"Reset your password",
		fmt.Sprintf("for test , token: %s", token),
	); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnInternalError, err)
		return
	}

	helpers.GenerateResponse(c, helpers.ReturnOK, nil)
	return
}

func ShowResetPage(c *gin.Context) {
	helpers.GenerateResponse(c, helpers.ReturnOK, nil)
	return
}

func ResetPassword(c *gin.Context) {
	helpers.GenerateResponse(c, helpers.ReturnOK, nil)
	return
}

func RefreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]
	db := helpers.GetDB(c)

	authToken, err := services.DecodeLoginToken(tokenString, db.Session(&gorm.Session{NewDB: true}))
	if err != nil {
		if !errors.Is(err, services.ErrorTokenExpired) {
			// todo: error handling
		}
	}

	if authToken.Revoked == false {
		// todo: error handling
	}

	// check if out of refresh limit
	refreshExpire := authToken.CreatedAt.Add(time.Duration(config.LoginAuth.RefreshExpire) * time.Second)
	if time.Now().After(refreshExpire) {
		// todo: error handling
	}

	// transaction
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer func() {
		if r := recover(); r != nil {
			log.Print(r.(error))
			tx.Rollback()
		}
	}()

	// generate token
	// delete old token
	// commit

	helpers.GenerateResponse(c, helpers.ReturnOK, nil)
	return
}
