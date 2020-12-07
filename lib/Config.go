package lib

import (
	"github.com/spf13/viper"
	"log"
	"path"
	"runtime"
)

var configPath = "/go/src/app"
var configName = "config"

func InitialConfigurations()  {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	//viper.SetDefault("database.dbname", "test_db")
}

func SetConfigName(name string) {
	configName = name
}

func SetAbsolutePath() {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		configPath = path.Dir(filename)
	}
}
