package centrifugoKit

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient([]string{"http://localhost:8000/api"}, "4e75aa67-3423-4e10-86b3-cfab4febef55", nil)
	if err != nil {
		panic(err)
	}

	m := map[string]interface{}{
		"msg": "hellO",
	}
	jsonStr, err := jsonKit.MarshalToString(m)
	if err != nil {
		panic(err)
	}

	rst, err := client.Publish(context.TODO(), "test-channel", []byte(jsonStr))
	if err != nil {
		panic(err)
	}
	fmt.Println("Epoch:", rst.Epoch)
	fmt.Println("Offset:", rst.Offset)
}
