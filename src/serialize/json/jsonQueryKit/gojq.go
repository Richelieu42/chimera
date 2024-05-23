package jsonQueryKit

import "github.com/itchyny/gojq"

var (
	// Parse 返回 *gojq.Query 实例
	Parse func(src string) (*gojq.Query, error) = gojq.Parse

	TypeOf func(v any) string = gojq.TypeOf

	Compare func(l, r any) int = gojq.Compare

	Compile func(q *gojq.Query, options ...gojq.CompilerOption) (*gojq.Code, error) = gojq.Compile

	Marshal func(v any) ([]byte, error) = gojq.Marshal

	NewIter func(values ...any) gojq.Iter = gojq.NewIter

	NewModuleLoader func(paths []string) gojq.ModuleLoader = gojq.NewModuleLoader
)
