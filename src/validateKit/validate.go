package validateKit

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

// New
/*
PS:
(1) 默认的tag name: "validate";
(2) Gin使用的tag name: "binding".

@param tagNameArgs 不传参的话，将采用默认的tag name
*/
func New(tagNameArgs ...string) *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())

	// 修改tag name
	tagName := sliceKit.GetFirstItemWithDefault("", tagNameArgs...)
	if strKit.IsNotEmpty(tagName) {
		v.SetTagName(tagName)
	}

	// 注册内置的验证器（自定义的）
	if err := registerBakedInValidation(v); err != nil {
		panic(err)
	}

	return v
}
