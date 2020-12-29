package handlers_v1

import (
	"auth/events"
	"auth/lib/event_listener"
	"auth/lib/helpers"
	"auth/models"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	db := c.MustGet("DB").(*gorm.DB)
	user, err := Register(input.Name, input.Email, input.Password, db)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyRegistered) {
			helpers.GenerateResponse(c, helpers.ReturnDuplicate, nil)
			return
		}

		helpers.GenerateResponse(c, helpers.ReturnInternalError, err.Error())
		return
	}

	newSession := db.Session(&gorm.Session{NewDB: true})
	token, err := GenerateActivationToken(user, newSession)
	if err != nil {
		helpers.GenerateResponse(c, helpers.ReturnInternalError, err.Error())
		return
	}

	dispatcher := c.MustGet("Dispatcher").(*event_listener.Dispatcher)
	dispatcher.Dispatch(events.EmailRegisteredEvent{User: *user, Token: token})

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

	var loginInfo models.EmailLogin
	db := c.MustGet("DB").(*gorm.DB)
	result := db.Where("email = ?", input.Email).First(&loginInfo)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		helpers.GenerateResponse(c, helpers.ReturnNotExist, nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(loginInfo.Password), []byte(input.Password)); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnNotExist, nil)
		return
	}

	// todo: produce jwt token for authentication

	helpers.GenerateResponse(c, helpers.ReturnOK, nil)
	return
}

func ActivateEmailRegister(c *gin.Context) {
	var input struct {
		Token string `form:"token"`
	}
	if err := c.ShouldBindQuery(&input); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, nil)
		return
	}

	db := c.MustGet("DB").(*gorm.DB)

	if err := Activate(input.Token, db); err != nil {
		if errors.Is(err, ErrorTokenNotValidYet) {
			helpers.GenerateResponse(c, helpers.ReturnTokenExpire, err)
			return
		} else if errors.Is(err, ErrorTokenExpired) {
			helpers.GenerateResponse(c, helpers.ReturnTokenExpire, err)
			return
		} else if errors.Is(err, ErrorTokenMalformed) {
			helpers.GenerateResponse(c, helpers.ReturnValidationFailed, err)
			return
		}

		helpers.GenerateResponse(c, helpers.ReturnInvalidToken, err)
		return
	}

	helpers.GenerateResponse(c, helpers.ReturnOK, nil)
	return
}
