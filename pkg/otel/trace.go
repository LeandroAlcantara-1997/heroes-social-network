package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	otl "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	trace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type config struct {
	serviceName string
	environment string
}

func New(serviceName, environment string) *config {
	return &config{
		serviceName: serviceName,
		environment: environment,
	}
}

func (c *config) TraceProvider(ctx context.Context) (*trace.TracerProvider, error) {
	exporter, err := otl.New(ctx)
	if err != nil {
		return nil, err
	}
	tc := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(c.serviceName),
			semconv.DeploymentEnvironmentKey.String(c.environment),
		)),
	)
	otel.SetTracerProvider(tc)
	return tc, nil
}
