package archiverKit

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestArchive(t *testing.T) {
	src := "/Users/richelieu/Documents/ino/images"
	mapper := map[string]string{
		src: "",
	}

	{
		out, err := os.Create(fmt.Sprintf("_%s.zip", idKit.NewXid()))
		if err != nil {
			logrus.Fatal(err)
		}
		defer out.Close()

		if err := ArchiveToZip(context.TODO(), out, mapper, nil); err != nil {
			logrus.Fatal(err)
		}
	}
	logrus.Info("---")

	{
		out, err := os.Create(fmt.Sprintf("_%s.tar.gz", idKit.NewXid()))
		if err != nil {
			logrus.Fatal(err)
		}
		defer out.Close()

		if err := ArchiveToTarGz(context.TODO(), out, mapper, nil); err != nil {
			logrus.Fatal(err)
		}
	}
	logrus.Info("---")
}
