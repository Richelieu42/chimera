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

	if isFileNameASCII(name) {
		w.Header().Set("Content-Disposition", `attachment; filename="`+name+`"`)
	} else {
		w.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(name))
	}

	http.ServeFile(w, r, path)
	return nil
}

func DownloadFileContent(w http.ResponseWriter, r *http.Request, content []byte, name string) error {
	if isFileNameASCII(name) {
		w.Header().Set("Content-Disposition", `attachment; filename="`+name+`"`)
	} else {
		w.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(name))
	}

	_, err := w.Write(content)
	return err
}

/*
https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
*/
func isFileNameASCII(fileName string) bool {
	for i := 0; i < len(fileName); i++ {
		if fileName[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
