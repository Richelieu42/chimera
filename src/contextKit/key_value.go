package contextKit

import "context"

// AttachKeyValue 在现有的上下文 (context) 中存储 1个 键值对，并返回一个新的上下文.
/*
PS:
(1) 同一个Context实例，多次调用此方法 且 key相同，则覆盖之前的值.
(2) 对于同一个Context实例，不推荐调用太多次此方法，因为: 每调用一次就是嵌套一层，会嵌套太多层.

@param ctx 不能为nil
*/
func AttachKeyValue[K comparable, V any](ctx context.Context, k K, v V) context.Context {
	return context.WithValue(ctx, k, v)
}

// AttachKeyValueMap 在现有的上下文 (context) 中存储 多个 键值对，并返回一个新的上下文.
/*
Deprecated: 如果传参m的元素数量多，可能会导致嵌套太多层，这样不太好.
			如果 键值对多 || 会频繁修改键值对，建议使用 gorilla/context（可能导致内存泄露!!!）.

@param ctx 不能为nil
*/
func AttachKeyValueMap[K comparable, V any](ctx context.Context, m map[K]V) context.Context {
	for k, v := range m {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}
