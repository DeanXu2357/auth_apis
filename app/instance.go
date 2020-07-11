package app

import (
    "sync"
	"fmt"
	"log"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Instance struct {
	Configs *Configurations
	Database *gorm.DB
}

var instance *Instance
var once sync.Once

func setup() (*Configurations, *gorm.DB) {
    InitConfig()
	config := GetConfigs()

    dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Database.DBHost,
		config.Database.DBPort,
		config.Database.DBUser,
		config.Database.DBName,
		config.Database.DBPassword)
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Printf("Database Connection failed : %s", err)
	}
	defer db.Close()

	db.DB().SetMaxIdleConns(4)
	db.DB().SetMaxOpenConns(8)
	db.DB().SetConnMaxLifetime(time.Minute)

    return config, db
}

func New() *Instance {
    once.Do(func() {
        c, d := setup()
        instance = &Instance{c, d}
    })
	return instance
}

func (i *Instance)GetConfig() *Configurations {
	return i.Configs
}
