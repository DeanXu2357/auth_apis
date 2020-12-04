package middlewares

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func SetDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)

		dbInfo := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			viper.Get("dbhost"),
			viper.Get("dbport"),
			viper.Get("user"),
			viper.Get("dbname"),
			viper.Get("dbpassword"))
		log.Printf("dbInfo: %s", dbInfo)

		db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
		if err != nil {
			log.Fatalf("Database Connection failed : %s", err)
		}

		c.Set("DB", db.WithContext(timeoutContext))
	}
}
