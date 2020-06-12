package controllers

import "github.com/gin-gonic/gin"

type registerByMailInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterByMail(c *gin.Context) {
	//
}

func ResendMail(c *gin.Context)  {
	//
}

func ActivateEmailRegister(c *gin.Context) {
	//
}
