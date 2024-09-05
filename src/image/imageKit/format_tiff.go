package imageKit

import (
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"golang.org/x/image/tiff"
	"os"
)

// ToTiff 将图片格式转换为".tiff".
func ToTiff(src, dest string, opts *tiff.Options) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	srcImage, _, err := DecodeWithImagePath(src)
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	return tiff.Encode(destFile, srcImage, opts)
}
