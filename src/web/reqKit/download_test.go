package reqKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"testing"
)

func TestDownloadToFile(t *testing.T) {
	url := "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"

	if err := DownloadToFile(url, "_file.png"); err != nil {
		panic(err)
	}
}

func TestDownloadToWriter(t *testing.T) {
	fileKit.NewFile("_writer.png")
}
