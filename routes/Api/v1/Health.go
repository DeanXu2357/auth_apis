package v1

import "github.com/gin-gonic/gin"

func Health(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
