package routes

import (
	v1 "auth/routes/Api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", v1.Health)

	return r
}
