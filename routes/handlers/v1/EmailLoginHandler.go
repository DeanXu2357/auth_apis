package handlers_v1

import (
	"auth/app"
	m "auth/models"
	"errors"
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

const EmailAlreadyRegistered = "email_already_registered"

func NewEmailController(app *app.Instance) *EmailLoginHandler {
	return &EmailLoginHandler{Application: app}
}

func (h *EmailLoginHandler)RegisterByMail(c *gin.Context) {
	var input registerByMailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 40022, "message": "validation failed"})
		return
	}

	name := input.Name
	email := input.Email
	password := input.Password

	user, err := h.Register(name, email, password);
	switch err.Error() {
	case EmailAlreadyRegistered:
		c.JSON(http.StatusBadRequest, gin.H{"status": 40009, "message": "email is already registered"})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "user_id": user.ID})
	return
}


func (h *EmailLoginHandler)ResendMail(c *gin.Context)  {
	//
}

func (h *EmailLoginHandler)ActivateEmailRegister(c *gin.Context) {
	//
}

func (h *EmailLoginHandler)Register(name string, email string, password string) (*m.User, error) {
	// todo : find a transaction manager library
	tx := h.Application.Database.Begin()

	user := &m.User{Name: name, Email: email}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, errors.New(EmailAlreadyRegistered)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return nil, err
	}

	if err := tx.Create(&m.EmailLogin{Email: email, Password: string(hashedPassword)}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()


	return user, nil
}
