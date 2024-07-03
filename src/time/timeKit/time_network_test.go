package timeKit

import (
	"context"
	"fmt"
	"testing"
)

func TestGetNetworkTime(t *testing.T) {
	fmt.Println(GetNetworkTime(context.TODO()))
}
