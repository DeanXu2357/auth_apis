package main

import (
	"auth/App"
	c "auth/Configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	configs := initConfig()
    app := App.App{
    	Configs: configs,
	}

	r := gin.Default()
	r.GET("/test", Test)
	r.Run(fmt.Sprintf(":%s", app.Configs.Server.port))
}

// Test 測試路由
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func initConfig() c.Configurations {
	viper.SetConfigName("app")
	viper.AddConfigPath("./Configs")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

    var config c.Configurations

	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configuration.Database.DBName)
	fmt.Println("Port is\t\t", configuration.Server.Port)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Database is\t", viper.GetString("database.dbname"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))

	return config
}
