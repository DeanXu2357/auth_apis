package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func SetDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeCtx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		c.Set("DB", db.Session(&gorm.Session{NewDB: true, Context: timeCtx}))
		c.Next()
	}
}
