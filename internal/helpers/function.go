package helpers

import (
	"auth/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var (
	NowFunction     = func() time.Time { return time.Now() }
	GetDBFunction   = func(c *gin.Context) *gorm.DB { return c.MustGet("DB").(*gorm.DB) }
	GetAuthFunction = func(c *gin.Context) models.AuthToken { return c.MustGet("Auth").(models.AuthToken) }
)

func NowTime() time.Time {
	return NowFunction()
}

func GetDB(c *gin.Context) *gorm.DB {
	return GetDBFunction(c)
}

func GetAuth(c *gin.Context) models.AuthToken {
	return GetAuthFunction(c)
}
