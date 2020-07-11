package controllers

import (
	"auth/app"
	"auth/services/email_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmailController struct {
	Application *app.Instance
	service *email_service.EmailService
}

type registerByMailInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewEmailController(app *app.Instance) *EmailController {
	service := email_service.New(app)
	return &EmailController{Application: app, service: service}
}

func (controller *EmailController)RegisterByMail(c *gin.Context) {
	var json registerByMailInput
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.service.RegistByMail(json.Email, json.Name, json.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
	return
}


func (controller *EmailController)ResendMail(c *gin.Context)  {
	//
}

func (controller *EmailController)ActivateEmailRegister(c *gin.Context) {
	//
}
