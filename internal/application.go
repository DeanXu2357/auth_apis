package application

import (
	"auth/lib/event_listener"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type Application struct {
	DB         *gorm.DB
	Dispatcher *event_listener.Dispatcher
	Tracer     opentracing.Tracer
}
