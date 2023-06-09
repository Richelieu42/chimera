package mailKit

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

func AssertEmail(email string) error {
	validate := validator.New()
	err := validate.Var(email, "required,email")
	if err != nil {
		return errorKit.NewSkip(1, "[%s] email(%s) is invalid with error(%s)", funcKit.GetFuncName(1), email, err.Error())
	}
	return nil
}
