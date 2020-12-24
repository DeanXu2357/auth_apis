package handlers_v1

import (
	"auth/events"
	"auth/lib/event_listener"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type registerByMailInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type verifyMailLogin struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const EmailAlreadyRegistered = "email_already_registered"

func RegisterByMail(c *gin.Context) {
	var input registerByMailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 40022, "message": "validation failed"})
		return
	}

	user, err := Register(input.Name, input.Email, input.Password, c.MustGet("DB").(*gorm.DB))
	if err != nil {
		switch err.Error() {
		case EmailAlreadyRegistered:
			c.JSON(http.StatusBadRequest, gin.H{"status": 40009, "message": "email is already registered"})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	dispatcher := c.MustGet("Dispatcher").(*event_listener.Dispatcher)
	dispatcher.Dispatch(events.NewEmailRegisteredEvent(*user))

	c.JSON(http.StatusOK, gin.H{"message": "success", "user_id": user.ID})
	return
}

func VerifyMailLogin(c *gin.Context) {
	var input registerByMailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 40022, "message": "validation failed"})
		return
	}

	//db := c.MustGet("DB").(*gorm.DB)

	//session :=
}

func ResendMail(c *gin.Context)  {
	//
}

func ActivateEmailRegister(c *gin.Context) {
	//
}
