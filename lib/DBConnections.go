package lib

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitialDatabase() *gorm.DB {
	timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)

	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		viper.GetString("db_host"),
		viper.GetString("db_port"),
		viper.GetString("db_user"),
		viper.GetString("db_name"),
		viper.GetString("db_password"))
	//log.Printf("dbInfo: %s", dbInfo)

	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database Connection failed : %s", err)
	}

	return db.WithContext(timeoutContext)
}
