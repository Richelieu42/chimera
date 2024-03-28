package otelKit

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

const (
	keyGinSpanCtx = "_chimera/gin-span-context"
	keyGinSpan    = "_chimera/gin-span"
)

// NewOuterGinMiddleware 适用于: （链路追踪）最外层的服务.
/*
PS: 需要先 set up.
*/
func NewOuterGinMiddleware() (middleware gin.HandlerFunc, err error) {
	if err = check(); err != nil {
		return
	}

	middleware = func(ctx *gin.Context) {
		tracer := otel.Tracer("")
		spanCtx, span := tracer.Start(ctx.Request.Context(), "entire")
		defer span.End()

		ctx.Set(keyGinSpanCtx, spanCtx)
		ctx.Set(keyGinSpan, span)

		ctx.Next()
	}
	return
}

// NewSecondaryGinMiddleware 适用于: （链路追踪）次级的服务.
/*
PS: 需要先 set up.
*/
func NewSecondaryGinMiddleware() (middleware gin.HandlerFunc, err error) {
	if err = check(); err != nil {
		return
	}

	middleware = func(ctx *gin.Context) {
		remoteSpanCtx, err := ExtractFromRequest(ctx.Request)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			ctx.Abort()
			return
		}

		tracer := otel.Tracer("")
		ctxWithSpan, span := tracer.Start(remoteSpanCtx, "entire")
		defer span.End()

		ctx.Set(keyGinSpanCtx, ctxWithSpan)
		ctx.Set(keyGinSpan, span)

		ctx.Next()
	}
	return
}

// GetSpanCtxFromGin
/*
!!!: 需要先使用中间件 NewOuterGinMiddleware 或 NewSecondaryGinMiddleware.
*/
func GetSpanCtxFromGin(ctx *gin.Context) context.Context {
	value, exists := ctx.Get(keyGinSpanCtx)
	if !exists || value == nil {
		return context.TODO()
	}
	return value.(context.Context)
}

// GetSpanFromGin
/*
!!!: 需要先使用中间件 NewOuterGinMiddleware 或 NewSecondaryGinMiddleware.
*/
func GetSpanFromGin(ctx *gin.Context) trace.Span {
	value, exists := ctx.Get(keyGinSpanCtx)
	if !exists || value == nil {
		tracer := NewNoopTracer()
		_, span := tracer.Start(context.TODO(), "noop")
		return span
	}
	return value.(trace.Span)
}
