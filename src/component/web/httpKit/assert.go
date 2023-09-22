package httpKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func AssertHttpUrl(httpUrl string) error {
	if err := validateKit.HttpUrl(httpUrl); err != nil {
		return errorKit.NewSkip(1, "[%s] httpUrl(%s) is because of with error(%s)", funcKit.GetFuncName(1), httpUrl, err.Error())
	}
	return nil
}
