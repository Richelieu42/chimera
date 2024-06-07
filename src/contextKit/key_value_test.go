package contextKit

import (
	"context"
	"fmt"
	"testing"
)

func TestWithKeyValueMap(t *testing.T) {
	ctx := WithKeyValueMap(context.TODO(), map[string]any{
		"a": "string",
		"b": true,
		"c": 666,
	})

	fmt.Println(ctx.Value("a").(string)) // string
	fmt.Println(ctx.Value("b").(bool))   // true
	fmt.Println(ctx.Value("c").(int))    // 666
}
