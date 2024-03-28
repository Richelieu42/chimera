package otelKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"time"
)

// MustSetUpWithGrpc
/*
@param opts e.g. otlptracegrpc.WithInsecure(), otlptracegrpc.WithDialOption(grpc.WithBlock())
*/
func MustSetUpWithGrpc(grpcEndpoint, serviceName string, attributeMap map[string]string, opts ...otlptracegrpc.Option) {
	err := SetUpWithGrpc(grpcEndpoint, serviceName, attributeMap, opts...)
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUpWithGrpc(grpcEndpoint, serviceName string, attributeMap map[string]string, opts ...otlptracegrpc.Option) error {
	if err := strKit.AssertNotEmpty(serviceName, "serviceName"); err != nil {
		return err
	}

	/* TracerProvider */
	tp, err := NewGrpcTracerProvider(grpcEndpoint, serviceName, attributeMap, opts...)
	if err != nil {
		return err
	}
	otel.SetTracerProvider(tp)

	/* TextMapPropagator */
	textMapPropagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	otel.SetTextMapPropagator(textMapPropagator)

	logrus.RegisterExitHandler(func() {
		ShutdownTracerProvider(tp, time.Second*3)
	})
	return nil
}
