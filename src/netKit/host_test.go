package netKit

import (
	"fmt"
	"testing"
)

func TestSplitHostnamePort(t *testing.T) {
	fmt.Println(SplitHost("localhost:8080")) // "localhost" "8080" <nil>

	fmt.Println(SplitHost("127.0.0.1")) // "" "" address 127.0.0.1: missing port in address
	fmt.Println(SplitHost("localhost")) // "" "" address localhost: missing port in address
}
