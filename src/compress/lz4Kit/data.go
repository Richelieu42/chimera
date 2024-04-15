package lz4Kit

import (
	"bytes"
	"github.com/pierrec/lz4/v4"
	"io"
)

func Compress(data []byte) (compressedData []byte, err error) {
	buf := bytes.NewBuffer(nil)

	lz4Writer := lz4.NewWriter(buf)
	_, err = lz4Writer.Write(data)
	if err != nil {
		return
	}
	// !!!: 需要先 关闭Writer 再调用 buf.Bytes().
	if err = lz4Writer.Close(); err != nil {
		return nil, err
	}

	compressedData = buf.Bytes()
	return
}

func Decompress(compressedData []byte) (decompressData []byte, err error) {
	lz4Reader := lz4.NewReader(bytes.NewReader(compressedData))
	decompressData, err = io.ReadAll(lz4Reader)
	return
}
