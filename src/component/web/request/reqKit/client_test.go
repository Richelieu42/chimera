package reqKit

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	data, err := client.Get("https://www.baidu.com").Do().ToBytes()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
