package slbKit

import (
	"testing"
)

func TestNewLoadBalancer(t *testing.T) {
	lb := NewLoadBalancer(nil)
	urls := []string{"http://127.0.0.1:8000", "http://127.0.0.1:8001", "http://127.0.0.1:8002"}
	for _, urlStr := range urls {
		backend, err := NewBackend(urlStr)
		if err != nil {
			panic(err)
		}
		if err := lb.AddBackend(backend); err != nil {
			panic(err)
		}
	}

}
