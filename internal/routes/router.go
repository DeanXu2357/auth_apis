package routes

import (
	handlerV1 "auth/internal/handlers/v1"
	"auth/internal/middlewares"
	"auth/lib/event_listener"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, d *event_listener.Dispatcher) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.SetDB(db))
	r.Use(middlewares.SetEventListener(d))

	routes := r.Group("/api")
	v1 := routes.Group("/v1")
	{
		v1.GET("/health", handlerV1.Health)

		emails := v1.Group("/email")
		emails.POST("/register", handlerV1.RegisterByMail)
		emails.POST("/verify", handlerV1.VerifyMailLogin)
		//emails.POST("/resend", handlerV1.ResendMail)
		emails.GET("/activate", handlerV1.ActivateEmailRegister)
	}

	return r
}
