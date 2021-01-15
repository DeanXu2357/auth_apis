package main

import (
	"auth/internal/cmd/migration"
	"auth/internal/cmd/sending_email"
	"auth/internal/config"
	"auth/internal/events"
	"auth/internal/listeners"
	"auth/internal/routes"
	"auth/lib/database"
	"auth/lib/email"
	"auth/lib/event_listener"
	log2 "auth/lib/log"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rootCmd := &cobra.Command{
		Short: "Console commands for this project",
	}

	config.InitialConfigurations()

	serveCmd := generateServerCmd()

	rootCmd.AddCommand(sending_email.GenerateCommand())
	rootCmd.AddCommand(generateTestCmd())
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(migration.GenerateCommand())

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func generateServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "run server",
		Run: func(cmd *cobra.Command, args []string) {
			db := database.InitialDatabase()

			dispatcher := event_listener.NewDispatcher()
			dispatcher.AttachListener(events.EmailRegistered, listeners.SendMailListener{})
			dispatcher.Consume()

			runServer(db, dispatcher)

			defer func() {
				log.Print("After shutdown server, close other objects")
				dispatcher.Close()
				sqlDB, _ := db.DB()
				err := sqlDB.Close()
				if err != nil {
					log.Fatal(err)
				}
			}()
		},
	}
}

func generateTestCmd() *cobra.Command {
	return &cobra.Command{
		Use: "test",
		Short: "for testing",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test success")

			dispatcher := event_listener.NewDispatcher()
			dispatcher.AttachListener(events.Test, listeners.PrintMsgListener{})
			defer dispatcher.Close()

			log.Print("do something")
			var e events.TestEvent
			dispatcher.Dispatch(e)
			log.Print("do rest of works")

			time.Sleep(5*time.Second)
			//testSendEmail()
		},
	}
}

func testSendEmail() {
	info := email.NewInfo()
	err := email.NewEmail(info).SendMail(
		[]string{"jasugun0000+receiver@gmail.com"},
		"test mail subject",
		"this is a test mail")
	if err != nil {
		fmt.Println(err)
	}
}

func runServer(db *gorm.DB, dispatcher *event_listener.Dispatcher) {
	router := routes.InitRouter(db, dispatcher)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", viper.Get("server_port")),
		Handler:      router,
		ReadTimeout:  30 * time.Second, // todo : be config
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log2.BeforeExit()
}
