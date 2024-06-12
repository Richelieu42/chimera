package httpKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"net/http"
)

// Redirect 请求重定向
/*
@param code [300, 308] || 201
*/
func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) error {
	if err := strKit.AssertNotEmpty(url, "url"); err != nil {
		return err
	}

	// Richelieu: 参考了 gin v1.10.0 中的 render.Redirect.Render()，对 code 进行了限制
	if (code < http.StatusMultipleChoices || code > http.StatusPermanentRedirect) && code != http.StatusCreated {
		return errorKit.Newf("can't redirect with status code(%d)", code)
	}

	http.Redirect(w, r, url, code)
	return nil
}
