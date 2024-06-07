package contextKit

import "context"

// WithKeyValueMap 在现有的上下文 (context) 中存储多个键值对，并返回一个新的上下文.
/*
@param ctx 不能为nil
*/
func WithKeyValueMap[K comparable, V any](ctx context.Context, m map[K]V) context.Context {
	for k, v := range m {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}
