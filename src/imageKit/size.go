package imageKit

import (
	"bytes"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

type (
	Size struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}
)

// GetSize 获取图片的尺寸（宽高）.
/*
【图像处理】Golang 获取常用图像的宽高总结
	https://www.cnblogs.com/voipman/p/16108320.html
*/
func GetSize(path string) (*Size, error) {
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	extName := fileKit.GetExtName(path)
	if err := strKit.AssertNotEmpty(extName, "extName"); err != nil {
		return nil, err
	}
	extName = strKit.ToLower(extName)

	var imgConf image.Config
	switch extName {
	case "jpg", "jpeg":
		imgConf, err = jpeg.DecodeConfig(bytes.NewReader(data))
	case "webp":
		imgConf, err = webp.DecodeConfig(bytes.NewReader(data))
	case "png":
		imgConf, err = png.DecodeConfig(bytes.NewReader(data))
	case "tif", "tiff":
		imgConf, err = tiff.DecodeConfig(bytes.NewReader(data))
	case "gif":
		imgConf, err = gif.DecodeConfig(bytes.NewReader(data))
	case "bmp":
		imgConf, err = bmp.DecodeConfig(bytes.NewReader(data))
	default:
		return nil, errorKit.New("invalid extName(%s)", extName)
	}
	if err != nil {
		return nil, err
	}
	return &Size{
		Width:  imgConf.Width,
		Height: imgConf.Height,
	}, nil
}
