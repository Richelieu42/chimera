package req111Kit

import (
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"testing"
)

func TestDownloadToFile(t *testing.T) {
	url := "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"

	if err := DownloadToFile(url, "_file.png"); err != nil {
		panic(err)
	}
}

func TestDownloadToWriter(t *testing.T) {
	url := "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"

	f, err := fileKit.Create("_writer.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := DownloadToWriter(url, f); err != nil {
		panic(err)
	}
}
