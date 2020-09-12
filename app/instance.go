package app

import (
	"log"
	"sync"
	"fmt"
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
	config := InitConfigs()

	db := initialDatabase(config)
	if db == nil {
		log.Fatal(db.Error)
	}

	return config, db
}

func initialDatabase(config *Configurations) (*gorm.DB) {
	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Database.DBHost,
		config.Database.DBPort,
		config.Database.DBUser,
		config.Database.DBName,
		config.Database.DBPassword)
	log.Printf("dbInfo: %s", dbInfo)

	// todo: retry and ping

	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalf("Database Connection failed : %s", err)
		return nil
	}

	db.DB().SetMaxIdleConns(4)
	db.DB().SetMaxOpenConns(8)
	db.DB().SetConnMaxLifetime(time.Minute)
	return db
}

func New() *Instance {
    once.Do(func() {
        instance = NewStatic()
    })
	return instance
}

func NewStatic() *Instance {
	c, d := setup()
	return &Instance{Configs: c, Database: d}
}

func (i *Instance)SetConfigs(c *Configurations) {
	i.Configs = c
}

func (i *Instance)GetConfig() *Configurations {
	return i.Configs
}
