package app

import "auth/configs"

type Instance struct {
	Configs *configs.Configurations
}

func New(c *configs.Configurations) *Instance {
	return &Instance{c}
}

func (i *Instance)GetConfig() *configs.Configurations {
	return i.Configs
}
