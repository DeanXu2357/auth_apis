package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"time"
)

func SetDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		span := opentracing.SpanFromContext(c.Request.Context())
		timeCtx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		ctx := opentracing.ContextWithSpan(timeCtx, span)

		c.Set("DB", db.Session(&gorm.Session{NewDB: true, Context: ctx}))
		c.Next()
	}
}
