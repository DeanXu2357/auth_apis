package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", Test)
	r.Run(":8080")
}

// Test 測試路由
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
