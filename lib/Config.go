package lib

import (
	"github.com/spf13/viper"
	"log"
)

func InitialConfigurations()  {
	viper.SetConfigName("config")
	viper.AddConfigPath("/go/src/app")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
}
