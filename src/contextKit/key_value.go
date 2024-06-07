package contextKit

import "context"

// WithKeyValueMap
/*
@param ctx 不能为nil
*/
func WithKeyValueMap(ctx context.Context, m map[any]any) context.Context {
	for k, v := range m {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}
