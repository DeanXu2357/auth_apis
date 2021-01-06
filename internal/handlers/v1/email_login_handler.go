package handlers_v1

import (
	"auth/internal/events"
	"auth/internal/helpers"
	"auth/lib/event_listener"
	"errors"
	"github.com/gin-gonic/gin"
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
	token, err := GenerateActivationToken(*user, newSession)
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
	token, err := EmailVerify(input.Email, input.Password, db)
	if err != nil {
		helpers.GenerateResponse(c, helpers.ReturnNotExist, nil)
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

func RecoveryPassword(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindQuery(&input); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, nil)
		return
	}
/*
	db := helpers.GetDB(c)

	// find login raw
	emailLogin, err := FindEmailLogin(input.Email, db)
	if err != nil {
		return
	}
	// genterate token and mail
	token, err := GeneratePasswordToken(emailLogin.User, db)
	if err != nil {
		// todo handle
	}
*/
	// mail

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
	helpers.GenerateResponse(c, helpers.ReturnOK, nil)
	return
}
