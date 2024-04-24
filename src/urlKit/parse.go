package urlKit

import "net/url"

// Parse 解析url（http、https、rtsp、rtmp等协议）
/*
PS:
(1) URL.EscapedPath():		处理过的 path
(2) URL.EscapedFragment():	处理过的 fragment（# 后面的玩意）
(3) Values.Encode():		处理过的 query string
(4) url.Parse vs url.ParseRequestURI: 当要解析的url字符串中包含有字符“#”时，使用url.Parse解析，会导致#后面的参数解析不出来。而使用ParseRequestURI就能解析到.
	使用场景:
		(a) rawURL中必定不存在'#'，使用:	url.Verify
		(b) rawURL中可能存在'#'，使用:		url.ParseRequestURI

GoLand教程-Go URL解析
	https://mp.weixin.qq.com/s/i6uEUzvu5BPna5QtSYDSMQ
golang url.Parse 解析
	https://blog.csdn.net/zhuyuqiang1238/article/details/121807708

@param rawURL !!!不能是如下格式: "localhost:8080"

e.g.
	u, err := urlKit.Verify("http://localhost:8080/go?a=123&b=456")
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)     // http
	fmt.Println(u.Host)       // localhost:8080
	fmt.Println(u.Hostname()) // localhost
	fmt.Println(u.Port())     // 8080
	fmt.Println(u.Path)       // /go
	fmt.Println(u.RawQuery)   // a=123&b=456
	fmt.Println(u.Query())    // map[a:[123] b:[456]]
*/
var Parse func(rawURL string) (*url.URL, error) = url.Parse

var ParseRequestURI func(rawURL string) (*url.URL, error) = url.ParseRequestURI
