package main

import (
	c "auth/configs"
	a "auth/app"
	"auth/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var application *a.Instance

func init() {
	c.InitConfig()
	config := c.GetConfigs()

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

	application = a.New(config, db)
}

func doAfterShutdown() {
	log.Print("doing events after shutdown")
}

func main() {
	router := routes.InitRouter(application)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", application.Configs.Server.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	doAfterShutdown()

	log.Println("Server exiting")
}
