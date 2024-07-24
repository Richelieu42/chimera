package contextKit

import (
	"context"
	"fmt"
	"testing"
)

func TestWithKeyValue(t *testing.T) {
	ctx := context.Background()

	ctx = WithKeyValue(ctx, "key", "a")
	fmt.Println("value:", ctx.Value("key")) // value: a

	ctx = WithKeyValue(ctx, "key", "b")
	ctx = WithKeyValue(ctx, "key", "c")
	fmt.Println("value:", ctx.Value("key")) // value: c
}

func TestWithKeyValue1(t *testing.T) {

}
