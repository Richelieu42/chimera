package timeKit

import (
	"context"
	"fmt"
	"testing"
)

func TestGetNetworkTime(t *testing.T) {
	for i := 0; i < 3; i++ {
		fmt.Println(GetNetworkTime(context.TODO()))
	}
}
