package routes

import (
	"auth/app"
	h "auth/routes/Api/v1"
	"auth/routes/Api/v1/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

var application *app.Instance

func InitRouter(app *app.Instance) *gin.Engine {
	r := gin.Default()
	application = app

	emailController := controllers.NewEmailController(app)

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
	dbName := &application.Configs.Database.DBName
	db := application.Database
	if dbName == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Undefined database name"})
		return
	}

	pingErr := db.DB().Ping()
	if pingErr != nil {
		c.AbortWithError(http.StatusInternalServerError, pingErr)
		return
	}

	c.JSON(200, gin.H{
		"message": dbName,
	})
}
