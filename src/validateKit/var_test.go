package validateKit

import (
	"fmt"
	"testing"
)

func TestHostname(t *testing.T) {
	fmt.Println(Hostname("www.yozo.com")) // <nil>
	fmt.Println(Hostname("10.0.9.141"))   // <nil>
}

func TestHost(t *testing.T) {
	fmt.Println(Host("www.yozo.com:8888")) // <nil>
	fmt.Println(Host("10.0.9.141:80"))     // <nil>

	// 非法端口: 0
	fmt.Println(Host("10.0.9.141:0")) // Key: '' Error:Field validation for '' failed on the 'hostname_port' tag
	// 非法端口: -1
	fmt.Println(Host("10.0.9.141:-1")) // Key: '' Error:Field validation for '' failed on the 'hostname_port' tag
}

func TestJson(t *testing.T) {
	fmt.Println(Json(""))             // Key: '' Error:Field validation for '' failed on the 'json' tag
	fmt.Println(Json("[}"))           // Key: '' Error:Field validation for '' failed on the 'json' tag
	fmt.Println(Json(`{name:123}`))   // Key: '' Error:Field validation for '' failed on the 'json' tag
	fmt.Println(Json(`{"name":123}`)) // <nil>
}
