package req111Kit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"io"
)

func DownloadToFile(url, filePath string) error {
	// url
	if err := httpKit.AssertHttpUrl(url); err != nil {
		return err
	}
	// filePath
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return err
	}

	client := GetDefaultClient()
	_, err := client.R().SetOutputFile(filePath).Get(url)
	if err != nil {
		return err
	}
	return nil
}

func DownloadToWriter(url string, writer io.Writer) error {
	// url
	if err := httpKit.AssertHttpUrl(url); err != nil {
		return err
	}
	// writer
	if err := interfaceKit.AssertNotNil(writer, "writer"); err != nil {
		return err
	}

	client := GetDefaultClient()
	_, err := client.R().SetOutput(writer).Get(url)
	if err != nil {
		return err
	}
	return nil
}
