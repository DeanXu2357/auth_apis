package handlers_v1

import (
	m "auth/models"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	name := input.Name
	email := input.Email
	password := input.Password
	db := c.MustGet("DB").(*gorm.DB)

	user, err := Register(name, email, password, db);
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

func Register(name string, email string, password string, db *gorm.DB) (*m.User, error) {
	// todo : find a transaction manager library
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Print(r.(error))
			tx.Rollback()
		}
	}()

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
