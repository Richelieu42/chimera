package ginKit

import "github.com/gin-gonic/gin"

type (
	ginOptions struct {
		// ServiceInfo 当前服务的信息（默认为""）
		/*
			e.g."Agent-127.0.0.1:12345"

			涉及: 	(a) 限流器
					(b) RespondError()、RespondPackage()、RespondPackageOrError()
		*/
		ServiceInfo string

		// RecoveryMiddleware panic恢复中间件（可以为nil，将采用默认值: gin.Recovery()）
		RecoveryMiddleware gin.HandlerFunc

		// DefaultOptionsMiddleware 是否使用 默认的Options中间件？
		DefaultOptionsMiddleware bool

		// DefaultNoRouteHtml 是否使用 默认的404页面？
		DefaultNoRouteHtml bool

		// DefaultNoMethod 是否使用 默认的405处理？
		DefaultNoMethod bool

		// DefaultFavicon 是否使用 默认的favicon.ico？
		DefaultFavicon bool
	}

	GinOption func(opts *ginOptions)
)

func loadOptions(options ...GinOption) *ginOptions {
	opts := &ginOptions{
		ServiceInfo:              "",
		RecoveryMiddleware:       nil,
		DefaultOptionsMiddleware: true,
		DefaultNoRouteHtml:       true,
		DefaultNoMethod:          true,
		DefaultFavicon:           true,
	}

	for _, option := range options {
		option(opts)
	}

	if opts.RecoveryMiddleware == nil {
		opts.RecoveryMiddleware = gin.Recovery()
	}

	return opts
}

func WithServiceInfo(serviceInfo string) GinOption {
	return func(opts *ginOptions) {
		opts.ServiceInfo = serviceInfo
	}
}

func WithRecoveryMiddleware(recoveryMiddleware gin.HandlerFunc) GinOption {
	return func(opts *ginOptions) {
		opts.RecoveryMiddleware = recoveryMiddleware
	}
}

func WithDefaultOptionsMiddleware(defaultOptionsMiddleware bool) GinOption {
	return func(opts *ginOptions) {
		opts.DefaultOptionsMiddleware = defaultOptionsMiddleware
	}
}

func WithDefaultNoRouteHtml(defaultNoRouteHtml bool) GinOption {
	return func(opts *ginOptions) {
		opts.DefaultNoRouteHtml = defaultNoRouteHtml
	}
}

func WithDefaultNoMethod(defaultNoMethod bool) GinOption {
	return func(opts *ginOptions) {
		opts.DefaultNoMethod = defaultNoMethod
	}
}

func WithDefaultFavicon(defaultFavicon bool) GinOption {
	return func(opts *ginOptions) {
		opts.DefaultFavicon = defaultFavicon
	}
}
