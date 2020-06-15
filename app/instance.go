package app

import (
	"auth/configs"
	"github.com/jinzhu/gorm"
)

type Instance struct {
	Configs *configs.Configurations
	Database *gorm.DB
}

func New(c *configs.Configurations, d *gorm.DB) *Instance {
	return &Instance{c, d}
}

func (i *Instance)GetConfig() *configs.Configurations {
	return i.Configs
}
