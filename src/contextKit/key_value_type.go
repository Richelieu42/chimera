package contextKit

import (
	"context"
)

type valueCtx struct {
	context.Context

	Data map[any]any
}

//func (c *valueCtx) String() string {
//	return contextName(c.Context) + ".WithValue(type " +
//		reflectlite.TypeOf(c.key).String() +
//		", val " + stringify(c.val) + ")"
//}

func (ctx *valueCtx) Value(key any) any {
	if value, ok := ctx.Data[key]; ok {
		return value
	}
	return ctx.Context.Value(key)
}
