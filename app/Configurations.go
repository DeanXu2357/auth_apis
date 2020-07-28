package app

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

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

var configPath = "./app"

func InitConfigs() *Configurations {
	var Configs *Configurations

	viper.SetConfigName("app")
	viper.AddConfigPath(configPath)
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

	return Configs
}

func SetConfigAbsolutePath() {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		configPath = path.Dir(filename)
	}
}
