package server

import (
	"auth/internal/events"
	"auth/internal/listeners"
	"auth/internal/routes"
	"auth/lib/database"
	"auth/lib/event_listener"
	log2 "auth/lib/log"
	myTracer "auth/lib/tracer"
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
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

type Application struct {
	DB         *gorm.DB
	dispatcher *event_listener.Dispatcher
	Tracer     opentracing.Tracer
}

func GenerateServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "run server",
		Run: func(cmd *cobra.Command, args []string) {
			db := database.InitialDatabase()

			dispatcher := event_listener.NewDispatcher()
			dispatcher.AttachListener(events.EmailRegistered, listeners.SendMailListener{})
			dispatcher.Consume()

			//tracer, tracerCloser, err := myTracer.NewJaegerTracer()
			//if err != nil {
			//	log.Fatal(err)
			//}

			runServer(db, dispatcher)

			defer func() {
				log.Print("After shutdown server, close other objects")
				dispatcher.Close()
				sqlDB, _ := db.DB()
				err := sqlDB.Close()
				if err != nil {
					log.Fatal(err)
				}
				tracerCloser.Close()
			}()
		},
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
