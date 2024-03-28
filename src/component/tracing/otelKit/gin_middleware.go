package otelKit

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// NewOuterGinMiddleware TODO: 适用于: （链路追踪）最外层的服务
/*
PS: 需要先 set up.
*/
func NewOuterGinMiddleware() (middleware gin.HandlerFunc, err error) {
	if err = check(); err != nil {
		return
	}

	middleware = func(ctx *gin.Context) {
		tracer := otel.Tracer("")
		ctxWithSpan, span := tracer.Start(ctx.Request.Context(), "main")
		defer span.End()

		ctx.Set(KeyGinContextWithSpan, ctxWithSpan)
		ctx.Set(KeyGinSpan, span)

		ctx.Next()
	}
	return
}

// NewSecondaryGinMiddleware TODO: 适用于: （链路追踪）次级的服务
/*
PS: 需要先 set up.
*/
func NewSecondaryGinMiddleware() (middleware gin.HandlerFunc, err error) {
	if err = check(); err != nil {
		return
	}

	middleware = func(ctx *gin.Context) {
		//tracer := otel.Tracer("", trace.WithInstrumentationAttributes(attribute.String("111", "111")))
		//spanCtx, span := tracer.Start(context.TODO(), "main", trace.WithAttributes(attribute.String("222", "222")))
		//defer span.End()
		//
		//ctx.Next()
	}
	return
}

func GetCtxWithSpanFromGinContext(ctx *gin.Context) context.Context {
	value, exists := ctx.Get(KeyGinContextWithSpan)
	if !exists {
		return context.TODO()
	}
	if value == nil {
		return context.TODO()
	}
	return value.(context.Context)
}

func GetSpanFromGinContext(ctx *gin.Context) trace.Span {
	value, exists := ctx.Get(KeyGinContextWithSpan)
	if !exists {
		return context.TODO()
	}
	if value == nil {
		return context.TODO()
	}
	return value.(trace.Span)
}
