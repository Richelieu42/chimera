package proxyKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/urlKit"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
	"log"
	"net/http"
	"net/http/httputil"
)

type (
	proxyOptions struct {
		ctx *gin.Context

		// scheme "http"（默认） || "https"
		scheme string

		// reqUrlPath 请求路由
		reqUrlPath *string

		// overrideQueryParams （优先级高于extraQueryParams）不会保留原先的queries
		overrideQueryParams map[string][]string

		// extraQueryParams 会保留原先的queries
		extraQueryParams map[string][]string

		// errorLogger 错误日志（可以为nil，但不建议这么干，因为错误会输出到控制台（通过 log.Printf()），不利于错误定位）
		errorLogger *log.Logger

		// polyfillHeaders 是否额外处理请求头?
		polyfillHeaders bool
	}

	ProxyOption func(opts *proxyOptions)
)

func loadOptions(options ...ProxyOption) *proxyOptions {
	opts := &proxyOptions{
		scheme:          "http",
		polyfillHeaders: true,
	}

	for _, option := range options {
		option(opts)
	}

	opts.scheme = strKit.EmptyToDefault(opts.scheme, "http", true)

	return opts
}

func (opts *proxyOptions) getClientIP(req *http.Request) string {
	if opts.ctx != nil {
		return opts.ctx.ClientIP()
	}
	return httpKit.GetClientIP(req)
}

// proxy
/*
@param errLogger 	可以为nil，但不建议这么干，因为错误会输出到控制台（通过 log.Printf()），不利于错误定位
@param scheme 		"http" || "https"
@param targetHost	e.g."127.0.0.1:8888"
@param reqUrlPath 	(1) 可以为nil（此时不修改 req.URL.Path）
					(2) 非nil的话，个人感觉: 字符串的第一个字符应该是"/"
@param extraQueryParams 	可以为nil
@return 可能的值:
		(1) context.Canceled
		(2) http.ErrAbortHandler，对应情况: 代理的 sse连接 或 http_stream连接(centrifugo) 断开时，会返回此error，可忽略

更多可参考:
httputil.NewSingleHostReverseProxy() https://m.bilibili.com/video/BV1H64y1u7D7?buvid=Y44D4D448DC195994A5A88CED2DA982C60DF&is_story_h5=false&mid=5%2BiuUUrTqJQOdIa1r3VR0g%3D%3D&p=1&plat_id=114&share_from=ugc&share_medium=iphone&share_plat=ios&share_session_id=8B36D2C9-4DCB-4BE5-80AD-F7D49E292B5F&share_source=WEIXIN&share_tag=s_i&timestamp=1680160438&unique_k=16bK0gz&up_id=456307879

PS:
(0) 通过 httputil.ReverseProxy 实现请求转发;
(1) 支持代理的协议: https、http、wss、ws...
(2) 如果请求转发的目标有效，但处理此请求需要花费大量时间（比如20+min），此时如果请求的客户端终端了请求（e.g.浏览器页面被直接关闭了），将返回 context.Canceled.
(3) targetHost有效，reqUrlPath非nil但事实上不存在该路由的情况，返回值为nil && 原始客户端得到404（404 page not found）.
(4) 代理请求前，如果读取了Request.Body的内容但不恢复（即重置其内容），将直接返回error（e.g.net/http: HTTP/1.x transport connection broken: http: ContentLength=161 with Body length 0）.
	且目标方不会收到请求.

e.g.	将 https://127.0.0.1:8888/test 转发给 http://127.0.0.1:8889/test
传参可以是：
(1) scheme=http targetHost=127.0.0.1:8889 reqUrlPath=nil
(2) scheme=http targetHost=127.0.0.1:8889 reqUrlPath=&"/test"
传参不能是：
(1) scheme=http targetHost=127.0.0.1:8889 reqUrlPath=&"test" （400 Bad Request）

e.g.1	将 https://127.0.0.1:8888/test 转发给 http://127.0.0.1:8889/test1
传参可以是：
(1) scheme=http targetHost=127.0.0.1:8889 reqUrlPath=&"/test1"
传参不能是：
(1) scheme=http targetHost=127.0.0.1:8889 reqUrlPath=&"test1"

e.g.2	将 https://127.0.0.1:8888/group/test 转发给 http://127.0.0.1:8889/test1
scheme="http" targetHost="127.0.0.1:8889" reqUrlPath=ptrKit.ToPtr("/test1")

e.g.3	将 https://127.0.0.1:8888/group/test 转发给 http://127.0.0.1:8889/group1/test1
scheme="http" targetHost="127.0.0.1:8889" reqUrlPath=ptrKit.ToPtr("/group1/test1")

e.g.4	将 wss://127.0.0.1:8888/test 转发给 ws://127.0.0.1:80/ws/connect
scheme="http" targetHost="127.0.0.1:80" reqUrlPath=ptrKit.ToPtr("/ws/connect")
*/
func (opts *proxyOptions) proxy(writer http.ResponseWriter, req *http.Request, targetHost string) (err error) {
	/* reset Request.Body */
	if err = httpKit.TryToResetRequestBody(req); err != nil {
		return
	}

	/* check targetHost */
	if err = validateKit.Var(targetHost, "hostname_port"); err != nil {
		err = errorKit.Wrapf(err, "invalid targetHost(%s)", targetHost)
		return
	}

	/* check scheme */
	scheme := opts.scheme
	switch scheme {
	case "https", "http":
	default:
		return errorKit.Newf("invalid scheme(%s)", scheme)
	}

	/* polyfill header */
	if opts.polyfillHeaders {
		/*
			(0) X-Forwarded-Proto: 客户端与代理服务器（或负载均衡服务器）间的连接所采用的传输协议（HTTP 或 HTTPS）
				!!!: 值不一定准确，除非 代理(s) 好好配合（有的话）.
		*/
		httpKit.SetHeader(req.Header, "X-Forwarded-Proto", httpKit.GetClientScheme(req))

		// (1) client ip
		tmp := httpKit.GetClientIPFromHeader(req)
		if strKit.IsEmpty(tmp) {
			httpKit.SetHeader(req.Header, "X-Real-IP", opts.getClientIP(req))
		}
	}

	/* Richelieu: 在请求头加个标记，证明此请求被 chimera 代理过 */
	mark(req.Header)

	/* proxy */
	director := func(req *http.Request) {
		req.URL.Scheme = opts.scheme
		req.URL.Host = targetHost
		if opts.reqUrlPath != nil {
			req.URL.Path = *opts.reqUrlPath
		}

		// 可能会修改 req.URL.RawQuery
		if opts.overrideQueryParams != nil {
			urlKit.OverrideRawQuery(req.URL, opts.overrideQueryParams)
		} else if opts.extraQueryParams != nil {
			urlKit.AddToRawQuery(req.URL, opts.extraQueryParams)
		}
	}
	reverseProxy := &httputil.ReverseProxy{
		Director: director,
		ErrorLog: opts.errorLogger,
		ErrorHandler: func(rw http.ResponseWriter, req *http.Request, e error) {
			err = errorKit.Wrapf(e, "fail to proxy")
		},
	}

	// Richelieu: 此处的 recover() 是针对 ReverseProxy.ServeHTTP() 中的 panic(http.ErrAbortHandler)
	defer func() {
		if obj := recover(); obj != nil {
			if err1, ok := obj.(error); ok {
				err = err1
				return
			}
			err = errorKit.Newf("recover from %v", obj)
		}
	}()
	reverseProxy.ServeHTTP(writer, req)

	return
}
