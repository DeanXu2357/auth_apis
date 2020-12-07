package middlewares

import (
	"auth/lib"
	"github.com/gin-gonic/gin"
)

func SetDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", lib.InitialDatabase())
	}
}
