package main

import (
	"auth/lib"
	"auth/routes"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func doAfterShutdown() {
	log.Print("doing events after shutdown")
}

func main() {
	lib.InitialConfigurations()
	db := lib.InitialDatabase()
	router := routes.InitRouter(db)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", viper.Get("server_port")),
		Handler: router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
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

	sqlDB, _ := db.DB()
	err := sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	doAfterShutdown()

	log.Println("Server exiting")
}
