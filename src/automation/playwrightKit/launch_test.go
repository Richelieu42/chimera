package playwrightKit

import (
	"github.com/playwright-community/playwright-go"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestLaunchBrowser(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		logrus.Infof("wd: [%s].", wd)
	}

	pw, browser, err := LaunchBrowser(BrowserNameChromium, "_playwright-deps", true, &playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		panic(err)
	}
	defer pw.Stop()
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		logrus.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://news.ycombinator.com"); err != nil {
		logrus.Fatalf("could not goto: %v", err)
	}

	time.Sleep(time.Second * 10)
}
