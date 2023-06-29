package ginKit

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// NewOpenCorsMiddleware 新建一个开放的（任意origin都通过）cors中间件.
func NewOpenCorsMiddleware() gin.HandlerFunc {
	return NewCorsMiddleware(nil)
}

// NewCorsMiddleware 新建一个cors中间件.
/*
「Go框架」 Gin 怎么实现允许前端跨域请求？
	https://mp.weixin.qq.com/s/2eJUJKJ3Xu5jOYXAfknqtA

@param origins 	(1) origin白名单（可以为nil）
				(2) 支持wildcard（*）
				(3) len(origins) == 0，则通配

e.g.
	传参: []string{"https://*.github.com", "https://api.*", "http://*", "https://facebook.com", "*.golang.org"}
*/
func NewCorsMiddleware(origins []string) gin.HandlerFunc {
	/*
		CORS 中间件提供三个函数，代表三种使用方式，分别是:
			cors.New()
			cors.Default()
			cors.DefaultConfig()
	*/

	config := cors.Config{
		// 允许的请求方式，默认值是 GET，POST，PUT，PATCH，DELETE，HEAD，和 OPTIONS
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		// 用在对预请求的响应中，指示实际的请求中可以使用哪些 HTTP 请求头
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		// 可以在响应中显示的请求头
		ExposeHeaders: []string{"Content-Length", "Content-Type", "Last-Modified"},
		// 预请求的结果能被缓存多久
		MaxAge: time.Hour * 12,

		// 添加请求源是否允许使用通配符，例如 http://some-domain/*，https://api. 或 http://some.*.subdomain.com
		AllowWildcard: true,
		// 表示请求附带请求凭据时是否响应请求，例如 cookie、HTTP authentication 或客户端 SSL 证书
		AllowCredentials: true,
		// 允许使用常用的浏览器的扩展模式
		AllowBrowserExtensions: true,
		// 允许使用 WebSocket 协议
		AllowWebSockets: true,
		// 允许使用 file:// 协议
		AllowFiles: true,
	}
	if len(origins) > 0 {
		/*
			(1) 允许部分origin
			AllowOrigins: 指定允许请求源的列表，如果列表中存在 *，则允许所有请求源，默认值是 []
		*/

		config.AllowOrigins = origins
	} else {
		/*
			(2) 允许全部origin
			AllowOriginFunc: 接收参数 origin，函数体中的验证逻辑返回是否允许跨域请求。该配置项优先级高于 AllowOrigins []string，如果设置该配置项，AllowOrigins []string 配置项的设置被忽略
		*/
		config.AllowOriginFunc = func(origin string) bool {
			return true
		}
	}
	return cors.New(config)
}
