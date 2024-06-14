package playwrightKit

import (
	"github.com/playwright-community/playwright-go"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestLaunchBrowser(t *testing.T) {
	_, browser, err := LaunchBrowser(BrowserNameChromium, "_temp", true, &playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	page, err := browser.NewPage()
	if err != nil {
		logrus.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://news.ycombinator.com"); err != nil {
		logrus.Fatalf("could not goto: %v", err)
	}

	time.Sleep(time.Second * 10)

}
