package log

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//go:generate mockgen -destination ../../mock/log_mock.go -package=mock -source=log.go
type Log interface {
	SendErrorLog(ctx context.Context, err error)
	SendEvent(ctx context.Context, eventCode int, message string)
}

type GlobalEvent struct {
	Data *EventMetadata `json:"event"`
}

type GlobalLog struct {
	Data *LogMetadata `json:"event"`
}
type EventMetadata struct {
	RequestID string `json:"requestId,omitempty"`
	Type      string `json:"type,omitempty"`
	DataType  string `json:"dataType,omitempty"`
	LogLevel  string `json:"logLevel,omitempty"`
	UserAgent string `json:"userAgent,omitempty"`
	EventCode int    `json:"eventCode,omitempty"`
	Message   string `json:"message,omitempty"`
}

type LogMetadata struct {
	LogLevel string  `json:"logLevel,omitempty"`
	Message  string  `json:"message,omitempty"`
	Request  Request `json:"request"`
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

const loggerKey string = "logger"

func AddLoggerInContext(ctx *gin.Context) {
	SetLogger(ctx)
}

func GetLogger(ctx context.Context) *GlobalLog {
	return ctx.Value(loggerKey).(*GlobalLog)
}
func SetLogger(ctx *gin.Context) {
	ctx.Set(loggerKey, &GlobalLog{
		Data: &LogMetadata{
			Request: Request{
				RequestID: uuid.NewString(),
				Host:      ctx.Request.Host,
				UserAgent: ctx.Request.Header.Get("User-Agent"),
				IP:        ctx.RemoteIP(),
				Method:    ctx.Request.Method,
				Endpoint:  ctx.Request.URL.Path,
				Headers:   ctx.Request.Header,
			},
		},
	})
}
