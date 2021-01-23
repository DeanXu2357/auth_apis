package database

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

const (
	callBackBeforeName = "opentracing:before"
	callBackAfterName  = "opentracing:after"
	gormSpanKey        = "span_key"
)

type OpentracingPlugin struct{}

func before(db *gorm.DB) {
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, "gorm")

	db.InstanceSet(gormSpanKey, span)
}

func after(db *gorm.DB) {
	s, isExist := db.InstanceGet(gormSpanKey)
	if !isExist {
		return
	}

	span, ok := s.(opentracing.Span)
	if !ok {
		return
	}
	defer span.Finish()

	if db.Error != nil {
		span.LogFields(log.Error(db.Error))
	}

	span.LogFields(log.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
	return
}

func (o *OpentracingPlugin) Name() string {
	return "open_tracing_integration"
}

func (o *OpentracingPlugin) Initialize(db *gorm.DB) error {
	db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return nil
}
