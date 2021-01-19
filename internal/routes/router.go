package routes

import (
	"auth/internal"
	handlerV1 "auth/internal/handlers/v1"
	"auth/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(s application.Application) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.SetDB(s.DB.Session(&gorm.Session{NewDB: true})))
	r.Use(middlewares.SetEventListener(s.Dispatcher))
	if s.Tracer != nil {
		r.Use(middlewares.Tracing(s.Tracer))
	}


	routes := r.Group("/api")
	v1 := routes.Group("/v1")
	{
		v1.GET("/health", handlerV1.Health)

		emails := v1.Group("/email")
		emails.POST("/register", handlerV1.RegisterByMail)
		emails.POST("/verify", handlerV1.VerifyMailLogin)
		emails.GET("/activate", handlerV1.ActivateEmailRegister)

		// 未完成
		//emails.POST("/recovery", handlerV1.RecoveryPassword)
		//emails.GET("/reset", handlerV1.ShowResetPage)
		//emails.POST("/reset", handlerV1.ResetPassword)

		user := v1.Group("/user")
		user.Use(middlewares.AuthorizeUserToken())
		user.GET("", handlerV1.ShowUser)

		v1.POST("refresh", handlerV1.RefreshToken)
	}

	return r
}
