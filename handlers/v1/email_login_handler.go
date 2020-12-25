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

type registerByMailInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type verifyMailLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const EmailAlreadyRegistered = "email_already_registered"

func RegisterByMail(c *gin.Context) {
	var input registerByMailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.GenerateResponse(c, helpers.ReturnValidationFailed, nil)
		return
	}

	user, err := Register(input.Name, input.Email, input.Password, c.MustGet("DB").(*gorm.DB))
	if err != nil {
		switch err.Error() {
		case EmailAlreadyRegistered:
			helpers.GenerateResponse(c, helpers.ReturnDuplicate, nil)
			return
		default:
			helpers.GenerateResponse(c, helpers.ReturnInternalError, err.Error())
			return
		}
	}

	dispatcher := c.MustGet("Dispatcher").(*event_listener.Dispatcher)
	dispatcher.Dispatch(events.NewEmailRegisteredEvent(*user))

	helpers.GenerateResponse(c,helpers.ReturnOK, map[string]interface{}{"user_id": user.ID, "email": input.Email})
	return
}

func VerifyMailLogin(c *gin.Context) {
	var input verifyMailLogin
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

	// todo: produce jwt token

	helpers.GenerateResponse(c, helpers.ReturnOK, nil)
}

func ActivateEmailRegister(c *gin.Context) {
	//
}
