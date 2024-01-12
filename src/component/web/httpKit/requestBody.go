package httpKit

import (
	"bytes"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
	"io"
	"net/http"
	"strings"
)

// MakeRequestBodySeekable
/*
Go语言: 如何让 request.Body 可以多次读取
	https://www.cnblogs.com/ayanmw/p/17191530.html

PS:
(1) 一般与 proxy() 搭配使用;
(2) 某个路由涉及代理（请求转发）的话，需要在handler里面 首先 调用此方法;
(3) 请求有内容的话，会全部读取一遍请求体内容.
*/
func MakeRequestBodySeekable(req *http.Request) error {
	// 特殊情况: req.Body == http.NoBody，http客户端发的是post请求，但是没有request body（即没post参数）
	if strKit.EqualsIgnoreCase(req.Method, http.MethodGet) || req.Body == nil || req.Body == http.NoBody {
		return nil
	}

	if _, ok := req.Body.(io.Seeker); ok {
		// 已经实现了 io.Seeker，避免重复调用
		return nil
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	// bytes.NewReader() 的返回值实现了 io.ReadSeeker 接口
	reader := bytes.NewReader(data)
	req.Body = ioKit.NopCloserToReadSeeker(reader)
	return nil
}

// ResetRequestBody 重置请求体，以防: 已经读完body了，请求转发给别人，别人收到的请求没内容.
/*
PS: req.Body可以为nil.
*/
func ResetRequestBody(req *http.Request) error {
	if strKit.EqualsIgnoreCase(req.Method, http.MethodGet) || req.Body == nil || req.Body == http.NoBody {
		// 无需重置
		return nil
	}

	seeker, ok := req.Body.(io.Seeker)
	if !ok {
		// 不能重置
		return errorKit.New("body(%T) is unable to seek", req.Body)
	}
	_, err := ioKit.SeekToStart(seeker)
	if err != nil {
		// 重置失败
		return err
	}
	return nil
}

// ToRequestBodyString
/*
条件:
(1) POST
(2) x-www-form-urlencoded

e.g.
	m := map[string][]string{
		"a": []string{"test"},
		"b": []string{"测试"},
	}
	fmt.Println(ToRequestBodyString(m)) // a=test&b=%E6%B5%8B%E8%AF%95
*/
func ToRequestBodyString(m map[string][]string) string {
	return urlKit.ToEscapedQueryString(m)
}

// OverrideRequestBody 覆盖request body.
/*
条件:
(1) POST
(2) x-www-form-urlencoded
*/
func OverrideRequestBody(req *http.Request, m map[string][]string) {
	content := ToRequestBodyString(m)
	reader := strings.NewReader(content)
	req.Body = ioKit.NopCloser(reader)
	req.ContentLength = int64(len(content))
}
