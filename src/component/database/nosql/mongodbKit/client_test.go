package mongodbKit

import (
	"context"
	"fmt"
	"testing"
)

func TestNewClientSimply(t *testing.T) {
	client, err := NewClientSimply(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(nil)

	{

	}
	fmt.Println("ok")
}
