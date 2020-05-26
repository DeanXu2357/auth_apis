package main

import (
	c "auth/configs"
	"auth/routes"
	"fmt"
)

func init() {
	c.InitConfig()
}

func main() {
	r := routes.InitRouter()
	_ = r.Run(fmt.Sprintf(":%v", c.Configs.Server.Port))
}
