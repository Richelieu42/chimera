package validateKit

import (
	"context"
	"github.com/go-playground/validator/v10"
)

// Struct 验证结构体.
/*
@param s 如果为nil，将返回error(e.g. validator: (nil *main.User))
*/
func Struct(s interface{}) error {
	v := New()
	return v.Struct(s)
}

func StructCtx(ctx context.Context, s interface{}) error {
	v := New()
	return v.StructCtx(ctx, s)
}

// StructExcept 验证结构体，排除指定的字段.
/*
@param fields 支持嵌套
*/
func StructExcept(s interface{}, fields ...string) error {
	v := New()
	return v.StructExcept(s, fields...)
}

func StructExceptCtx(ctx context.Context, s interface{}, fields ...string) error {
	v := New()
	return v.StructExceptCtx(ctx, s, fields...)
}

// StructPartial 验证结构体，只验证指定的字段.
/*
@param fields 支持嵌套
*/
func StructPartial(s interface{}, fields ...string) error {
	v := New()
	return v.StructPartial(s, fields...)
}

func StructPartialCtx(ctx context.Context, s interface{}, fields ...string) error {
	v := New()
	return v.StructPartialCtx(ctx, s, fields...)
}

// StructFiltered 验证结构体，过滤指定的字段.
func StructFiltered(s interface{}, fn validator.FilterFunc) error {
	v := New()
	return v.StructFiltered(s, fn)
}

func StructFilteredCtx(ctx context.Context, s interface{}, fn validator.FilterFunc) error {
	v := New()
	return v.StructFilteredCtx(ctx, s, fn)
}
