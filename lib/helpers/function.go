package helpers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var Helper HelperInstance

type HelperInstance struct{}

func (h HelperInstance) getDB(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}

func (h HelperInstance) nowTime() time.Time {
	return time.Now()
}

func init() {
	Helper = HelperInstance{}
}

func NowTime() time.Time {
	return Helper.nowTime()
}

func GetDB(c *gin.Context) *gorm.DB {
	return Helper.getDB(c)
}
