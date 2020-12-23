package main

import (
	"auth/cmd/sending_email"
	"auth/events"
	"auth/lib"
	"auth/lib/email"
	"auth/lib/event_listener"
	"auth/listeners"
	"auth/routes"
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

	lib.InitialConfigurations()

	serveCmd := &cobra.Command{
		Use: "serve",
		Short: "run server",
		Run: func(cmd *cobra.Command, args []string) {
			db := lib.InitialDatabase()

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

	rootCmd.AddCommand(sending_email.GenerateCommand())
	rootCmd.AddCommand(generateTestCmd())
	rootCmd.AddCommand(serveCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
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
			dispatcher.Consume()
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
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
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
}
