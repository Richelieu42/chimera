package contextKit

import "context"

// WithKeyValue 在现有的上下文 (context) 中存储 1个 键值对，并返回一个新的上下文.
func WithKeyValue[K comparable, V any](ctx context.Context, k K, v V) context.Context {
	return context.WithValue(ctx, k, v)
}

// WithKeyValueMap 在现有的上下文 (context) 中存储 多个 键值对，并返回一个新的上下文.
/*
@param ctx 不能为nil
*/
func WithKeyValueMap[K comparable, V any](ctx context.Context, m map[K]V) context.Context {
	for k, v := range m {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}
