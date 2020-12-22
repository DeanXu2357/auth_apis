package main

import (
	"auth/cmd/sending_email"
	"auth/lib"
	"auth/lib/email"
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
			runServer(db)
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

			testSendEmail()
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

func runServer(db *gorm.DB) {
	router := routes.InitRouter(db)

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

	log.Println("Server exiting")
}
