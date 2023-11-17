// Package ginKit
// 路由相关
package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
)

// RegisterHandlers
/*
适用场景: 1个路由，n个Method.

@param methods nil => 接收所有类型method的请求.	e.g. http.MethodGet、http.MethodPost
*/
func RegisterHandlers(group IGroup, route string, methods []string, handlers ...gin.HandlerFunc) error {
	if len(handlers) == 0 {
		return nil
	}
	sliceKit.ForEach(handlers, func(handler gin.HandlerFunc, index int) {

	},
	)

	if len(methods) == 0 {
		// (1) Any
		group.Any(route, handlers...)
	} else {
		// (2) 指定类型的method
		for _, method := range methods {
			group.Handle(method, route, handlers...)
		}
	}
	return nil
}

// RegisterHandlersRoutes 将多个相同的处理器，注册到多个路由.
func RegisterHandlersRoutes(group IGroup, routes []string, methods []string, handlers ...gin.HandlerFunc) (err error) {
	for _, route := range routes {
		err = RegisterHandlers(group, route, methods, handlers...)
		if err != nil {
			return
		}
	}
	return
}
