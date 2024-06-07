package contextKit

import "golang.org/x/net/context"

// WithKeyValue 在现有的上下文 (context) 中存储 1个 键值对，并返回一个新的上下文.
/*
@param ctx 不能为nil
*/
func WithKeyValue[K comparable, V any](ctx context.Context, k K, v V) context.Context {
	return context.WithValue(ctx, k, v)
}

// WithKeyValueMap 在现有的上下文 (context) 中存储 多个 键值对，并返回一个新的上下文.
/*
Deprecated: 如果传参m的元素数量多，可能会导致嵌套太多层，这样不太好.
			如果 键值对多 || 会频繁修改键值对，建议使用 gorilla/context.

@param ctx 不能为nil
*/
func WithKeyValueMap[K comparable, V any](ctx context.Context, m map[K]V) context.Context {
	for k, v := range m {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}
