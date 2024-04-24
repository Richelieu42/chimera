package validateKit

type (
	// Validatable Deprecated: 纯粹用 github.com/go-playground/validator/v10 吧.
	/*
		参考: "github.com/ardielle/ardielle-go/rdl"中的 rdl.Validatable
		!!!: 如果结构体实现了 Validatable 接口，方法体内部不能调用 validateKit.Struct，以免发生递归死循环（但可以调用 validateKit.Sign）.
	*/
	Validatable interface {
		Validate() error
	}
)
