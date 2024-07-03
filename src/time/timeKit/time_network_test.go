package timeKit

import (
	"fmt"
	"testing"
)

func TestGetNetworkTime(t *testing.T) {
	for i := 0; i < 3; i++ {
		fmt.Println(GetNetworkTime())
	}
}
