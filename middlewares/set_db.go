package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
	}
}
