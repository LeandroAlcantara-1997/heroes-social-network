package log

import (
	"context"
)

//go:generate mockgen -destination ../../../mock/log_mock.go -package=mock -source=log.go
type Log interface {
	SendErrorLog(ctx context.Context, message string)
	SendEvent(ctx context.Context, eventCode int, message string)
}

type GenericsData interface {
	EventMetadata | LogMetadata
}

type GlobalEvent struct {
	Data EventMetadata `json:"event"`
}

type GlobalLog struct {
	Data LogMetadata `json:"event"`
}
type EventMetadata struct {
	RequestID string `json:"requestId,omitempty"`
	Type      string `json:"type,omitempty"`
	DataType  string `json:"dataType,omitempty"`
	LogLevel  string `json:"logLevel,omitempty"`
	UserAgent string `json:"userAgent,omitempty"`
	EntityID  string `json:"entityId,omitempty"`
	EventCode int    `json:"eventCode,omitempty"`
	Message   string `json:"message,omitempty"`
}

type LogMetadata struct {
	RequestID string `json:"requestId,omitempty"`
	EntityID  string `json:"entityId,omitempty"`
	LogLevel  string `json:"logLevel,omitempty"`
	Type      string `json:"type,omitempty"`
	EventCode string `json:"eventCode,omitempty"`
	Message   string `json:"message,omitempty"`
	UserAgent string `json:"userAgent,omitempty"`
}
