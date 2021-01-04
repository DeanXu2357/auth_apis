package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	LoginAuth     LoginAuthSettings
	ActivateAuth  ActivateAuthSettings
	EventListener EventListenerSettings
)

func init() {
	InitialConfigurations()

	if err := viper.UnmarshalKey("activate_auth", &ActivateAuth); err != nil {
		log.Fatal(err)
	}

	if err := viper.UnmarshalKey("login_auth", &LoginAuth); err != nil {
		log.Fatal(err)
	}

	if err := viper.UnmarshalKey("event_listener", &EventListener); err != nil {
		log.Fatal(err)
	}
}

func InitialConfigurations() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/go/src/app")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
}
