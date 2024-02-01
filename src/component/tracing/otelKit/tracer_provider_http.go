package otelKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

// NewHttpTracerProvider 通过otlptracehttp
/*
PS: 此方法第二个返回值为nil的情况下，建议立即调用 otel.SetTracerProvider.

@param endpoint 	(1) 可以为""，将采用默认值: "localhost:4318"
					(2) e.g."localhost:4318"
@param serviceName 	服务名
@param attributeMap	可以为nil
@param opts 		额外配置（不要涉及 otlptracehttp.WithEndpoint，因为在此处配了也没用）
*/
func NewHttpTracerProvider(endpoint, serviceName string, attributeMap map[string]string, opts ...otlptracehttp.Option) (*trace.TracerProvider, error) {
	// 默认: 使用非安全协议（http）
	opts = append([]otlptracehttp.Option{otlptracehttp.WithInsecure()}, opts...)
	// 放在最后面（优先级最高）
	if strKit.IsNotEmpty(endpoint) {
		opts = append(opts, otlptracehttp.WithEndpoint(endpoint))
	}
	// 创建 exporter 实例
	exporter, err := otlptracehttp.New(context.TODO(), opts...)
	if err != nil {
		return nil, err
	}

	res := NewResourceWithAttributes(serviceName, attributeMap)

	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exporter),
		// Record information about this application in a Resource.
		trace.WithResource(res),
	)
	return tp, nil
}
