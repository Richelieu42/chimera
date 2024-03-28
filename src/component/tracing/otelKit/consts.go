package otelKit

const (
	// HeaderBaggage 跨服务相关的请求头
	HeaderBaggage = "Baggage"

	KeyTraceId = "trace-id"
	KeySpanId  = "span-id"

	KeyGinContextWithSpan = "_chimera/gin-context-with-span"
	KeyGinSpan            = "_chimera/gin-span"
)
