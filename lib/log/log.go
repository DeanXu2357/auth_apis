package log

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

type LogType string

const (
	TypeAccess LogType = "access_log"
	TypeSystem LogType = "system_log"
)

var (
	Logger *zap.Logger
)

func init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
}

func NewLogger() *zap.Logger {
	return Logger
}

func BeforeExit() {
	err := Logger.Sync()
	if err != nil {
		log.Fatalf("can't sync zap logger: %v", err)
	}
}

func Info(msg string, identify interface{}) {
	Log(zapcore.InfoLevel, msg, identify)
}

func Error(msg string, identify interface{}) {
	Log(zapcore.ErrorLevel, msg, identify)
}

func Warn(msg string, identify interface{})  {
	Log(zapcore.WarnLevel, msg, identify)
}

func Debug(msg string, identify interface{}) {
	Log(zapcore.DebugLevel, msg, identify)
}

func Log(level zapcore.Level, msg string, identify interface{}) {
	c, OK := identify.(*gin.Context)
	if OK {
		if ce := Logger.Check(level, msg); ce != nil {
			headers := c.Request.Header.Clone()
			headerBytes, _ := json.Marshal(headers)

			ce.Write(
				zap.Time("time", time.Now()),
				zap.String("log_type", string(TypeAccess)),
				zap.String("task_id", ""),
				zap.Int("status", c.Writer.Status()),
				zap.String("client_ip", c.ClientIP()),
				zap.String("http_method", c.Request.Method),
				zap.String("url", c.Request.URL.Path),
				zap.String("url_params", c.Request.URL.Query().Encode()),
				//zap.String("http_body", c.Reque)
				zap.String("http_header", string(headerBytes)),
			)
		}
	} else {
		if ce := Logger.Check(level, msg); ce != nil {
			ce.Write(
				zap.Time("time", time.Now()),
				zap.String("log_type", identify.(string)),
				zap.String("task_id", ""),
			)
		}
	}
}
