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
//	// Uncompress 解压缩.
//	/*
//	   PS: 大小限制: 100MB.
//	*/
//	Uncompress func(bs []byte) ([]byte, error) = codec.Gunzip
//)

var (
	Compress func(data []byte, level ...int) ([]byte, error) = gcompress.Gzip

	Uncompress func(data []byte) ([]byte, error) = gcompress.UnGzip
)

var (
	GzipFile func(srcFilePath, dstFilePath string, level ...int) (err error) = gcompress.GzipFile

	GzipPathWriter func(filePath string, writer io.Writer, level ...int) error = gcompress.GzipPathWriter
)

var (
	UnGzipFile func(srcFilePath, dstFilePath string) error = gcompress.UnGzipFile
)
