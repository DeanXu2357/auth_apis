package routes

import (
	"auth/configs"
	v1 "auth/routes/Api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", v1.Health)
	r.GET("/test_db", ShowDB)

	return r
}

func ShowDB(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": configs.Configs.Database.DBName,
	})
}
