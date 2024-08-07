package slbKit

import (
	"testing"
)

func TestNewLoadBalancer(t *testing.T) {
	lb := NewLoadBalancer(nil)
	lb.AddBackend()

}
