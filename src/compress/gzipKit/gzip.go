package gzipKit

import (
	"github.com/gogf/gf/v2/encoding/gcompress"
	"io"
)

//import "github.com/zeromicro/go-zero/core/codec"
//
//var (
//	// Compress 压缩.
//	Compress func(bs []byte) []byte = codec.Gzip
//
//	// Decompress 解压缩.
//	/*
//	   PS: 大小限制: 100MB.
//	*/
//	Decompress func(bs []byte) ([]byte, error) = codec.Gunzip
//)

// Compress
/*
PS: 不涉及 compressThreshold 的话，建议直接使用 Compress.
*/
func Compress(data []byte, options ...GzipOption) ([]byte, error) {
	opts := loadOptions(options...)

	return opts.Compress(data)
}

var (
	//Compress func(data []byte, level ...int) ([]byte, error) = gcompress.Gzip

	Decompress func(data []byte) ([]byte, error) = gcompress.UnGzip
)

var (
	GzipFile func(srcFilePath, dstFilePath string, level ...int) (err error) = gcompress.GzipFile

	GzipPathWriter func(filePath string, writer io.Writer, level ...int) error = gcompress.GzipPathWriter
)

var (
	UnGzipFile func(srcFilePath, dstFilePath string) error = gcompress.UnGzipFile
)
