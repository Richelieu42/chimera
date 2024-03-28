package otelKit

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// NewTracer
/*
PS:
(1) 需要先 set up;
(2) 使用全局的TracerProvider.
*/
func NewTracer(name string, opts ...trace.TracerOption) trace.Tracer {
	return otel.Tracer(name, opts...)
}
