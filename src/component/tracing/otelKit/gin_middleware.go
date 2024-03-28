package otelKit

import (
	"github.com/gin-gonic/gin"
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
		//tracer := otel.Tracer("", trace.WithInstrumentationAttributes(attribute.String("111", "111")))
		//spanCtx, span := tracer.Start(context.TODO(), "main", trace.WithAttributes(attribute.String("222", "222")))
		//defer span.End()
		//
		//ctx.Next()
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
