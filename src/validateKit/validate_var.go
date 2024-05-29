package validateKit

import "context"

// Var 验证变量.
func Var(field interface{}, tag string) error {
	v := New()
	return v.Var(field, tag)
}

func VarCtx(ctx context.Context, field interface{}, tag string) error {
	v := New()
	return v.VarCtx(ctx, field, tag)
}

// VarWithValue 用于验证一个值是否满足另一个值的某种条件.
/*
@param field	要验证的值
@param other	要与之比较的值
@param tag		一个表示验证规则的字符串
*/
func VarWithValue(field interface{}, other interface{}, tag string) error {
	v := New()
	return v.VarWithValue(field, other, tag)
}

func VarWithValueCtx(ctx context.Context, field interface{}, other interface{}, tag string) error {
	v := New()
	return v.VarWithValueCtx(ctx, field, other, tag)
}
