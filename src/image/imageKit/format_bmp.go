package imageKit

import (
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"golang.org/x/image/bmp"
	"os"
)

// ToBmp 将图片格式转换为".bmp".
func ToBmp(src, dest string) error {
	if err := fileKit.AssertNotExistOrIsFile(dest, true); err != nil {
		return err
	}

	srcImage, _, err := DecodeWithPath(src)
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	return bmp.Encode(destFile, srcImage)
}
