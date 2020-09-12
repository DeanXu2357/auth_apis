package routes

import (
	"auth/app"
	handlerV1 "auth/routes/handlers/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

var application *app.Instance

func InitRouter(app *app.Instance) *gin.Engine {
	r := gin.Default()
	application = app

	emailLoginHandler := handlerV1.NewEmailController(app)

	r.GET("/test_db", ShowDB)
	routes := r.Group("/api")
	v1 := routes.Group("/v1")
	{
		v1.GET("/health", handlerV1.Health)

		emails := v1.Group("/email")
		emails.POST("/register", emailLoginHandler.RegisterByMail)
		emails.POST("/resend", emailLoginHandler.ResendMail)
		emails.POST("/activate", emailLoginHandler.ActivateEmailRegister)
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
