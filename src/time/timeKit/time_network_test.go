package timeKit

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetNetworkTime(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	fmt.Println(GetNetworkTime(ctx))

	//select {}
}
