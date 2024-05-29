package validateKit

import "context"

// ValidateMap 验证map.
func ValidateMap(data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	v := New()
	return v.ValidateMap(data, rules)
}

func ValidateMapCtx(ctx context.Context, data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	v := New()
	return v.ValidateMapCtx(ctx, data, rules)
}
