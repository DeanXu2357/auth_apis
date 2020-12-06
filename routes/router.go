package routes

import (
	handlerV1 "auth/handlers/v1"
	"auth/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.SetDB())

	routes := r.Group("/api")
	v1 := routes.Group("/v1")
	{
		v1.GET("/health", handlerV1.Health)

		emails := v1.Group("/email")
		emails.POST("/register", handlerV1.RegisterByMail)
		emails.POST("/resend", handlerV1.ResendMail)
		emails.POST("/activate", handlerV1.ActivateEmailRegister)
	}

	return r
}
