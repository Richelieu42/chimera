package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	cancel()

	GetNetworkTime(ctx)

	select {}
}
