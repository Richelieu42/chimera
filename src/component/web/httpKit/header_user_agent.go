package httpKit

import (
	"github.com/medama-io/go-useragent"
	"net/http"
)

// GetUserAgent 获取http请求头中"User Agent"的值.
/*
参考: https://www.sunzhongwei.com/golang-gin-for-user-agent-in-http-request-header-value

e.g.
Chrome浏览器: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36
Safari浏览器: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.5 Safari/605.1.15
*/
func GetUserAgent(header http.Header) string {
	return GetHeader(header, "User-Agent")
}

// GetUserAgentInfo
/*
亚微秒级的解析速度！Go 语言的高性能 User-Agent 解析库
	https://mp.weixin.qq.com/s/Qvj64_WofoCZ8j0eTCq6cQ

PS: 如果是 Postman 发的http请求，返回值的所有字段都是零值.
*/
func GetUserAgentInfo(header http.Header) *useragent.UserAgent {
	str := GetUserAgent(header)
	tmp := useragent.NewParser().Parse(str)
	return &tmp
}
