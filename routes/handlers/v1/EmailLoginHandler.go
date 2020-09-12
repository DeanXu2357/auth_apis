package handlers_v1

import (
	"auth/app"
	m "auth/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

type EmailLoginHandler struct {
	Application *app.Instance
}

type registerByMailInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewEmailController(app *app.Instance) *EmailLoginHandler {
	return &EmailLoginHandler{Application: app}
}

func (controller *EmailLoginHandler)RegisterByMail(c *gin.Context) {
	var input registerByMailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 40022, "message": "validation failed"})
		return
	}

	tx := controller.Application.Database.Begin()

	if err := tx.Create(&m.User{Name: input.Name, Email: input.Email}).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 40009, "message": "email is already registered"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tx.Create(&m.EmailLogin{Email: input.Email, Password: string(hashedPassword)}).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "success"})
	return
}


func (controller *EmailLoginHandler)ResendMail(c *gin.Context)  {
	//
}

func (controller *EmailLoginHandler)ActivateEmailRegister(c *gin.Context) {
	//
}

func hashPassword(pwd string) string {
	return pwd
}

