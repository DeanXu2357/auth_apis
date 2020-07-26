package routes

import (
	"auth/app"
	h "auth/routes/Api/v1"
	"auth/routes/Api/v1/controllers"
	"github.com/gin-gonic/gin"
)

var application *app.Instance

func InitRouter(application *app.Instance) *gin.Engine {
	r := gin.Default()

	emailController := controllers.NewEmailController(application)

	r.GET("/test_db", ShowDB)
	routes := r.Group("/api")
	v1 := routes.Group("/v1")
	{
		v1.GET("/health", h.Health)

		emails := v1.Group("/email")
		emails.POST("/register", emailController.RegisterByMail)
		emails.POST("/resend", emailController.ResendMail)
		emails.POST("/activate", emailController.ActivateEmailRegister)
	}

	return r
}

func ShowDB(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": string(application.Configs.Database.DBName),
	})
}
