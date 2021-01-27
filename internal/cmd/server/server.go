package server

import (
	"auth/docs"
	"auth/internal"
	"auth/internal/config"
	"auth/internal/events"
	"auth/internal/listeners"
	"auth/internal/routes"
	"auth/lib/database"
	"auth/lib/event_listener"
	log2 "auth/lib/log"
	myTracer "auth/lib/tracer"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func GenerateServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "run server",
		Run: func(cmd *cobra.Command, args []string) {
			db := database.NewDBEngine()
			db.Use(&database.OpentracingPlugin{})

			dispatcher := event_listener.NewDispatcher()
			dispatcher.AttachListener(events.EmailRegistered, listeners.SendMailListener{})
			dispatcher.Consume()

			tracer, tracerCloser, err := myTracer.NewJaegerTracer(
				viper.GetString("app_name"),
				fmt.Sprintf("%s:%d", config.Tracer.AgentHost, config.Tracer.AgentPort),
				fmt.Sprintf("%s:%d", config.Tracer.SamplerHost, config.Tracer.SamplerPort),
			)
			if err != nil {
				log.Fatal(err)
			}

			runServer(application.Application{DB: db, Dispatcher: dispatcher, Tracer: tracer})

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

func runServer(s application.Application) {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = config.Swagger.Title
	docs.SwaggerInfo.Description = config.Swagger.Description
	docs.SwaggerInfo.Version = config.Swagger.Version
	docs.SwaggerInfo.Host = config.Swagger.Host
	docs.SwaggerInfo.BasePath = config.Swagger.BasePath
	docs.SwaggerInfo.Schemes = []string{"http"}

	router := routes.InitRouter(s)

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log2.BeforeExit()
}
