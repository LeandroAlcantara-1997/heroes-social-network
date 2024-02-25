package log

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

//go:generate mockgen -destination ../../mock/log_mock.go -package=mock -source=log.go
type Logger interface {
	Error(ctx context.Context, err error, data any)
}

type Vendor interface {
	Send(ctx context.Context, payload []byte)
}

type logger struct {
	Data   *Metadata `json:"event"`
	stdout *zap.Logger
	vendor Vendor
}

func NewLogger(environment string, vendor Vendor, zapLogger *zap.Logger) *logger {
	return &logger{
		stdout: zapLogger,
		vendor: vendor,
	}
}

type logLevel string

const logLevelError logLevel = "error"

type Metadata struct {
	LogLevel logLevel `json:"logLevel,omitempty"`
	Data     any      `json:"data,omitempty"`
	Message  string   `json:"message,omitempty"`
	Request  Request  `json:"request"`
}

type Request struct {
	RequestID string      `json:"requestId,omitempty"`
	Host      string      `json:"host"`
	UserAgent string      `json:"userAgent,omitempty"`
	IP        string      `json:"ip"`
	Method    string      `json:"method"`
	Endpoint  string      `json:"endpoint"`
	Headers   http.Header `json:"headers"`
}
type contextKey string

const loggerKey contextKey = "logger"

func GetLoggerFromContext(ctx context.Context) Logger {
	return ctx.Value(loggerKey).(Logger)
}

func (l *logger) Error(ctx context.Context, err error, data any) {
	l.Data = &Metadata{
		LogLevel: logLevelError,
		Message:  err.Error(),
		Data:     data,
		Request:  l.Data.Request,
	}

	l.stdout.Error(err.Error(), zap.String("message", l.Data.Message),
		zap.Any("data", l.Data.Data))
	payload, err := json.Marshal(&l)
	if err != nil {
		l.stdout.Error(err.Error(), zap.Any("data", l.Data))
	}
	l.vendor.Send(ctx, payload)
}

func (l *logger) NewLoggerMiddleware(ctx *gin.Context) {
	l.Data = &Metadata{
		Request: Request{
			RequestID: uuid.NewString(),
			Host:      ctx.Request.Host,
			UserAgent: ctx.Request.Header.Get("User-Agent"),
			IP:        ctx.RemoteIP(),
			Method:    ctx.Request.Method,
			Endpoint:  ctx.Request.URL.Path,
			Headers:   ctx.Request.Header,
		},
	}
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), loggerKey, l))
}

func AddLoggerInContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}
