package helpers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var NowFunction = func() time.Time {return time.Now()}
var GetDBFunction = func(c *gin.Context) *gorm.DB {return c.MustGet("DB").(*gorm.DB)}

func NowTime() time.Time {
	return NowFunction()
}

func GetDB(c *gin.Context) *gorm.DB {
	return GetDBFunction(c)
}
