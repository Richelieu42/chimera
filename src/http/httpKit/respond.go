package httpKit

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/jsonKit"
	"net/http"
	"net/url"
	"unicode"
)

const (
	PlainContentType = "text/plain; charset=utf-8"
	JsonContentType  = "application/json; charset=utf-8"
)

// Status 设置响应的http状态码
func Status(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

// RespondString
/*
参考: gin里面的 Context.String() .
*/
func RespondString(w http.ResponseWriter, code int, format string, values ...any) error {
	data := strKit.StringToBytes(fmt.Sprintf(format, values...))
	return RespondData(w, code, PlainContentType, data)
}

// RespondJson
/*
参考: gin里面的 Context.JSON() .
*/
func RespondJson(w http.ResponseWriter, code int, obj any) error {
	data, err := jsonKit.Marshal(obj)
	if err != nil {
		return err
	}
	return RespondData(w, code, JsonContentType, data)
}

// RespondFile 响应文件
/*
参考: gin里面的 Context.File() 和 Context.FileAttachment() .

@param filePath 文件路径
@param fileName 文件名（可以为""，可以与 传参filePath 里的不同）
@return 如果不为nil，建议输出到控制台
*/
func RespondFile(w http.ResponseWriter, r *http.Request, code int, filePath, fileName string) error {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return err
	}

	Status(w, code)
	if strKit.IsNotEmpty(fileName) {
		// https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
		isASCII := func(s string) bool {
			for i := 0; i < len(s); i++ {
				if s[i] > unicode.MaxASCII {
					return false
				}
			}
			return true
		}
		if isASCII(fileName) {
			w.Header().Set("Content-Disposition", `attachment; filename="`+fileName+`"`)
		} else {
			w.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(fileName))
		}
	}
	http.ServeFile(w, r, filePath)
	return nil
}

// RespondData 响应字节流（二进制流）
/*
参考: gin里面的 Context.Data() .

@return 如果不为nil，建议输出到控制台
*/
func RespondData(w http.ResponseWriter, code int, contentType string, data []byte) error {
	Status(w, code)

	writeContentType(w, []string{contentType})

	if !bodyAllowedForStatus(code) {
		return nil
	}
	_, err := w.Write(data)
	return err
}

// writeContentType
/*
PS:
(1) copy from gin/render/render.go
*/
func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
/*
PS:
(1) copy from gin/context.go
(2) @return 在对应http状态码的情况下，是否允许写内容？
*/
func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}
	return true
}
