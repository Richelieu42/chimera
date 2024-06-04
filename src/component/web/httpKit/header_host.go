package httpKit

import "net/http"

// SetHost 修改请求头中的Host，一般用于请求转发前.
func SetHost(req *http.Request, newHost string) {
	/*
		Richelieu: 下面的代码无效，参考:
		如何正确在 Golang 中在处理 Http Request 之前修改 Host 字段内容
			https://mp.weixin.qq.com/s/_hkJm2rT9SfaOGZwUh197Q
	*/
	//req.Header.Set("Host", newHost)

	req.Host = newHost
}

func GetHost(req *http.Request) string {
	/* Richelieu: 下面两种方法都不行 */
	//return req.URL.Host
	//return req.Header.Get("Host")

	return req.Host
}
