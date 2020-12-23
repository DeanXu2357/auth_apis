package middlewares

import (
	"auth/lib/event_listener"
	"github.com/gin-gonic/gin"
)

func SetEventListener(d *event_listener.Dispatcher) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Dispatcher", d)
		c.Next()
	}
}
