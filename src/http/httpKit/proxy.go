package httpKit

import (
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/urlKit"
	"log"
	"net/http"
	"net/http/httputil"
)

// Proxy 代理请求（反向代理，请求转发）
/*
@param errLogger 	可以为nil，但不建议这么干，因为错误会输出到控制台（通过 log.Printf()），不利于错误定位
@param scheme 		"http" || "https"
@param addr 		e.g."127.0.0.1:8888"
@param reqUrlPath 	(1) 可以为nil（此时不修改 req.URL.Path）
					(2) 非nil的话，个人感觉: 字符串的第一个字符应该是"/"
@param extraQuery 	可以为nil
@return 可能是 context.Canceled（可以用==进行比较）

更多可参考:
httputil.NewSingleHostReverseProxy() https://m.bilibili.com/video/BV1H64y1u7D7?buvid=Y44D4D448DC195994A5A88CED2DA982C60DF&is_story_h5=false&mid=5%2BiuUUrTqJQOdIa1r3VR0g%3D%3D&p=1&plat_id=114&share_from=ugc&share_medium=iphone&share_plat=ios&share_session_id=8B36D2C9-4DCB-4BE5-80AD-F7D49E292B5F&share_source=WEIXIN&share_tag=s_i&timestamp=1680160438&unique_k=16bK0gz&up_id=456307879

PS:
(1) 支持代理的协议: https、http、wss、ws...
(2) 如果请求转发的目标有效，但处理此请求需要花费大量时间（比如20+min），此时如果请求的客户端终端了请求（e.g.浏览器页面被直接关闭了），将返回 context.Canceled.
(3) addr有效，reqUrlPath非nil但事实上不存在该路由的情况，返回值为nil && 原始客户端得到404（404 page not found）.

e.g.	将 https://127.0.0.1:8888/test 转发给 http://127.0.0.1:8889/test
传参可以是：
(1) scheme=http addr=127.0.0.1:8889 reqUrlPath=nil
(2) scheme=http addr=127.0.0.1:8889 reqUrlPath=&"/test"
传参不能是：
(1) scheme=http addr=127.0.0.1:8889 reqUrlPath=&"test" （400 Bad Request）

e.g.1	将 https://127.0.0.1:8888/test 转发给 http://127.0.0.1:8889/test1
传参可以是：
(1) scheme=http addr=127.0.0.1:8889 reqUrlPath=&"/test1"
传参不能是：
(1) scheme=http addr=127.0.0.1:8889 reqUrlPath=&"test1"

e.g.2	将 https://127.0.0.1:8888/group/test 转发给 http://127.0.0.1:8889/test1
scheme="http" addr="127.0.0.1:8889" reqUrlPath=strKit.GetStringPtr("/test1")

e.g.3	将 https://127.0.0.1:8888/group/test 转发给 http://127.0.0.1:8889/group1/test1
scheme="http" addr="127.0.0.1:8889" reqUrlPath=strKit.GetStringPtr("/group1/test1")
*/
func Proxy(w http.ResponseWriter, r *http.Request, errorLogger *log.Logger, scheme, addr string, reqUrlPath *string, extraQuery map[string]string) (err error) {
	scheme = strKit.EmptyToDefault(scheme, "http", true)
	switch scheme {
	case "https":
	case "http":
	default:
		return errorKit.Simple("invalid scheme: %s", scheme)
	}
	if strKit.IsEmpty(addr) {
		return errorKit.Simple("addr is empty")
	}

	director := func(req *http.Request) {
		req.URL.Scheme = scheme
		req.URL.Host = addr
		if reqUrlPath != nil {
			req.URL.Path = *reqUrlPath
		}
		req.URL.RawQuery = urlKit.CombineQueryString(req.URL.RawQuery, urlKit.ToQueryString(extraQuery))
	}
	proxy := &httputil.ReverseProxy{
		Director: director,
		ErrorLog: errorLogger,
		ErrorHandler: func(rw http.ResponseWriter, req *http.Request, e error) {
			err = e
		},
	}
	proxy.ServeHTTP(w, r)
	return
}
