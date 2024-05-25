package netKit

import (
	"net"
	"strconv"
)

var (
	// SplitHostnamePort
	/*
		e.g.
			fmt.Println(SplitHostnamePort("localhost:8080")) // "localhost" "8080" <nil>
			fmt.Println(SplitHostnamePort("127.0.0.1")) // "" "" address 127.0.0.1: missing port in address
			fmt.Println(SplitHostnamePort("localhost")) // "" "" address localhost: missing port in address
	*/
	SplitHostnamePort func(host string) (hostname, port string, err error) = net.SplitHostPort
)

// JoinHostnameAndPort
/*
e.g.
	fmt.Println(netKit.JoinHostnameAndPort("127.0.0.1", 80)) // 127.0.0.1:80
	fmt.Println(netKit.JoinHostnameAndPort("", 8888))        // :8888
*/
func JoinHostnameAndPort(hostname string, port int) string {
	//return fmt.Sprintf("%s:%d", hostname, port)

	return net.JoinHostPort(hostname, strconv.Itoa(port))
}
