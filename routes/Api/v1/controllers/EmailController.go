package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerByMailInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterByMail(c *gin.Context) {
	var json registerByMailInput
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func ResendMail(c *gin.Context)  {
	//
}

func ActivateEmailRegister(c *gin.Context) {
	//
}
