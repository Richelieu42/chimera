package fileKit

import (
	"github.com/gabriel-vasile/mimetype"
	"io"
	"net/http"
)

// DetectContentType 获取 ContentType(即MimeType).
/*
PS: 读取前 512 个字节.

@return 保底 "application/octet-stream"

e.g.
([]byte(nil))	=> "text/plain; charset=utf-8"
([]byte{}) 		=> "text/plain; charset=utf-8"
*/
func DetectContentType(data []byte) string {
	return http.DetectContentType(data)
}

// Detect
/*
PS:
(1) mimetype库: 基于magic数的用于媒体类型和文件扩展名检测的快速的 Go 库，支持 170+ 格式.
(2) 读取前 3072 个字节.

e.g.
	mime := mimeTypeKit.Detect(nil)
	fmt.Println(mime.ToDSN()) // "text/plain"
*/
func Detect(in []byte) *mimetype.MIME {
	return mimetype.Detect(in)
}

func DetectReader(r io.Reader) (*mimetype.MIME, error) {
	return mimetype.DetectReader(r)
}

// DetectFile
/*
PS: 默认limit为: 3KB（3072）.

TODO: https://github.com/gabriel-vasile/mimetype
	mimetype.SetLimit(1024*1024) // Set limit to 1MB.
	// or
	mimetype.SetLimit(0) // No limit, whole file content used.
	mimetype.DetectFile("file.doc")

e.g.
	mime, _ := mimeTypeKit.DetectFile("/Users/richelieu/Desktop/未命名.wps")
	fmt.Println(mime.ToDSN()) // application/x-ole-storage

	mime, _ = mimeTypeKit.DetectFile("/Users/richelieu/Desktop/download.pdf")
	fmt.Println(mime.ToDSN()) // application/pdf
*/
func DetectFile(path string) (*mimetype.MIME, error) {
	return mimetype.DetectFile(path)
}
