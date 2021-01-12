package middlewares

import (
	"auth/internal/helpers"
	"auth/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthorizeUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len("Bearer "):]

		db := helpers.GetDB(c)

		authToken, err := services.DecodeLoginToken(tokenString, db.Session(&gorm.Session{NewDB: true}))
		if err != nil {
			helpers.GenerateResponse(c, helpers.ReturnBadRequest, map[string]string{"detail": err.Error()})
			c.Abort()
		}

		c.Set("Auth", authToken)
		c.Next()
	}
}
