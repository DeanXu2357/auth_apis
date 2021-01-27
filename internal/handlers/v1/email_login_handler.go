package handlers_v1

import (
	"auth/internal/events"
	"auth/internal/helpers"
	"auth/internal/models"
	"auth/internal/services"
	"auth/lib/email"
	"auth/lib/event_listener"
	log2 "auth/lib/log"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterByMailInput struct {
	Name     string `json:"name" binding:"required" example:"dean"`
	Email    string `json:"email" binding:"required" example:"dean.test@gmail.com"`
	Password string `json:"password" binding:"required" example:"!AS$GK())"`
}

// RegisterByMail godoc
// @Summary Register an account by email
// @Description Using email to register an account
// @Tags E-mail
// @Accept  json
// @Produce  json
// @Param JSON body RegisterByMailInput true "User data"
// @Success 200 {object} helpers.ResponseContent string "{"status":200, "msg":ok}"
// @Failure 400 {object} helpers.ResponseContent string "40022:validation failed, 400009: already registered"
// @Failure 500
// @Router /api/v1/email/register [post]
func RegisterByMail(c *gin.Context) {
	var input RegisterByMailInput
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

	log2.Info(fmt.Sprintf("activate token : %s", token), c)
	dispatcher := c.MustGet("Dispatcher").(*event_listener.Dispatcher)
	dispatcher.DispatchAsync(events.EmailRegisteredEvent{User: *user, Token: token})

	helpers.GenerateResponse(c, helpers.ReturnOK, map[string]interface{}{"user_id": user.ID, "email": input.Email})
	return
}

type VerifyMailLoginInput struct {
	Email    string `json:"email" binding:"required" example:"dean.test@gmail.com"`
	Password string `json:"password" binding:"required" example:"!AS$GK())"`
}

// VerifyMailLogin godoc
// @Summary Get login token by email
// @Description Using email to receive a login token
// @Tags E-mail
// @Accept  json
// @Produce  json
// @Param JSON body VerifyMailLoginInput true "login data"
// @Success 200 {object} helpers.ResponseContent string "{"status":200, "msg":ok}"
// @Failure 400 {object} helpers.ResponseContent string "40004:user not found, 40000: email not verified yet"
// @Failure 500
// @Router /api/v1/email/verify [post]
func VerifyMailLogin(c *gin.Context) {
	var input VerifyMailLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, nil)
		return
	}

	db := helpers.GetDB(c)
	token, err := services.EmailVerify(input.Email, input.Password, db)
	if err != nil {
		if errors.Is(err, services.ErrorUserNotFound) {
			helpers.GenerateResponse(c, helpers.ReturnNotExist, map[string]string{"detail": err.Error()})
			return
		}

		helpers.GenerateResponse(c, helpers.ReturnBadRequest, map[string]string{"detail": err.Error()})
		return
	}

	helpers.GenerateResponse(c, helpers.ReturnOK, map[string]string{"token": token})
	return
}

type ActivateEmailInput struct {
	Token string `form:"token"`
}

// ActivateEmail godoc
// @Summary activate email
// @Description activate email
// @Tags E-mail
// @Accept  json
// @Produce  json
// @Param token query ActivateEmailInput true "email authentication token"
// @Success 200 {object} helpers.ResponseContent string "{"status":200, "msg":ok}"
// @Failure 400 {object} helpers.ResponseContent string "40022:validation failed , 40102: token expired, 40101: unknown token invalid error"
// @Failure 500
// @Router /api/v1/email/activate [get]
func ActivateEmail(c *gin.Context) {
	var input ActivateEmailInput
	if err := c.ShouldBindQuery(&input); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, map[string]string{"detail": err.Error()})
		return
	}

	db := helpers.GetDB(c)

	if err := services.Activate(input.Token, db); err != nil {
		if errors.Is(err, services.ErrorTokenNotValidYet) {
			helpers.GenerateResponse(c, helpers.ReturnTokenExpire, map[string]string{"detail": err.Error()})
			return
		} else if errors.Is(err, services.ErrorTokenExpired) {
			helpers.GenerateResponse(c, helpers.ReturnTokenExpire, map[string]string{"detail": err.Error()})
			return
		} else if errors.Is(err, services.ErrorTokenMalformed) {
			helpers.GenerateResponse(c, helpers.ReturnValidationFailed, map[string]string{"detail": err.Error()})
			return
		}

		helpers.GenerateResponse(c, helpers.ReturnInvalidToken, map[string]string{"detail": err.Error()})
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
