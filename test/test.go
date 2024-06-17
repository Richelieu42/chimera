package main

import (
	"github.com/playwright-community/playwright-go"
	"github.com/richelieu-yang/chimera/v3/src/automation/playwrightKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	url := "https://www.moulem.com/"

	pw, browser, err := playwrightKit.LaunchBrowser(playwrightKit.BrowserNameChromium, "_playwright-deps", true, &playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		panic(err)
	}
	defer pw.Stop()
	defer browser.Close()
	bctx, err := browser.NewContext()
	if err != nil {
		panic(err)
	}
	page, err := bctx.NewPage()
	if err != nil {
		panic(err)
	}
	if _, err = page.Goto(url); err != nil {
		panic(err)
	}

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 6)
	logrus.Info("sleep ends")

	{
		data, err := page.Screenshot(playwright.PageScreenshotOptions{
			Path:     playwright.String("_test.png"),
			Type:     playwright.ScreenshotTypePng,
			FullPage: playwright.Bool(true),
		})
		if err != nil {
			panic(err)
		}
		logrus.Info(len(data))
	}

	//{
	//	locator := page.Locator("input#search")
	//	count, err := locator.Count()
	//	if err != nil {
	//		panic(err)
	//	}
	//	logrus.Infof("count: %d", count)
	//	if err := locator.Fill("hello world!"); err != nil {
	//		panic(err)
	//	}
	//}
	//
	//{
	//	locator := page.Locator("input#searchBtn")
	//	count, err := locator.Count()
	//	if err != nil {
	//		panic(err)
	//	}
	//	logrus.Infof("count: %d", count)
	//	if err := locator.Click(); err != nil {
	//		panic(err)
	//	}
	//}

	//resp, err := page.Reload()
	//if err != nil {
	//	panic(err)
	//}
	//logrus.Infof("reload: %t", resp.Ok())

	select {}
}
