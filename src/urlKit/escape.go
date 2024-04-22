package urlKit

import "net/url"

// EncodeURIComponent 编码.
/*
!!!: 事实上，"不完全"等同于js的 encodeURI || encodeURIComponent.

e.g.	golang
("") => ""
("test 测试") => "test+%E6%B5%8B%E8%AF%95"

e.g.1	js
encodeURIComponent("test 测试") => "test%20%E6%B5%8B%E8%AF%95"
*/
var EncodeURIComponent func(s string) string = url.QueryEscape

// DecodeURIComponent 解码.
/*
e.g.
("") => "", nil
*/
var DecodeURIComponent func(s string) (string, error) = url.QueryUnescape
