package netKit

import (
	"net"
	"strconv"
)

var (
	// SplitHost
	/*
		e.g.
			fmt.Println(SplitHost("localhost:8080")) // "localhost" "8080" <nil>
			fmt.Println(SplitHost("127.0.0.1")) // "" "" address 127.0.0.1: missing port in address
			fmt.Println(SplitHost("localhost")) // "" "" address localhost: missing port in address
	*/
	SplitHost func(host string) (hostname, port string, err error) = net.SplitHostPort
)

// JoinToHost
/*
e.g.
	fmt.Println(netKit.JoinToHost("127.0.0.1", 80)) // 127.0.0.1:80
	fmt.Println(netKit.JoinToHost("", 8888))        // :8888
*/
func JoinToHost(hostname string, port int) string {
	//return fmt.Sprintf("%s:%d", hostname, port)

	return net.JoinHostPort(hostname, strconv.Itoa(port))
}
