package httpKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"net/http"
	"net/url"
	"unicode"
)

func DownloadFile(w http.ResponseWriter, r *http.Request, path, name string) error {
	if strKit.IsEmpty(name) {
		name = fileKit.GetFileName(path)
	}

	// https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
	isASCII := func(s string) bool {
		for i := 0; i < len(s); i++ {
			if s[i] > unicode.MaxASCII {
				return false
			}
		}
		return true
	}
	if isASCII(name) {
		w.Header().Set("Content-Disposition", `attachment; filename="`+name+`"`)
	} else {
		w.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(name))
	}

	http.ServeFile(w, r, path)
	return nil
}
