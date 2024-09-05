package imageKit

import (
	"bytes"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"image"
	"io"
	"os"
)

// Decode 解码图片（部分特殊格式不支持）.
/*
@param r 类型可以是: *os.File（用完记得调用Close()）
@return 第1个: image.Image实例
		第2个: 表示图像的格式名称，例如 "png"、"jpeg" 等（不带"." && 转为小写）
		第3个: error（可能为nil）
*/
var Decode func(r io.Reader) (img image.Image, format string, err error) = image.Decode

// DecodeWithImagePath 解码图片.
/*
@param imagePath 图片的路径.
*/
func DecodeWithImagePath(imagePath string) (image.Image, string, error) {
	if err := fileKit.AssertExistAndIsFile(imagePath); err != nil {
		return nil, "", err
	}

	f, err := os.Open(imagePath)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	return Decode(f)
}

// DecodeWithBytes []byte => image.Image
func DecodeWithBytes(imageData []byte) (image.Image, string, error) {
	if err := sliceKit.AssertNotEmpty(imageData, "imageData"); err != nil {
		return nil, "", err
	}

	// 将 []byte 数据转换为 io.Reader
	imgReader := bytes.NewReader(imageData)

	return Decode(imgReader)
}
