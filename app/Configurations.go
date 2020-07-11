package app

import (
	"fmt"
	"github.com/spf13/viper"
)

var Configs *Configurations

type Configurations struct {
	Server ServerConfigurations
	Database DatabaseConfigurations
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPassword string
}

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./configs/")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&Configs)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}

func GetConfigs() (configs *Configurations) {
	 configs = Configs
	 return
}
