package main

import (
	"bufio"
	"github.com/playwright-community/playwright-go"
	"github.com/richelieu-yang/chimera/v3/src/automation/playwrightKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	url = "https://61.160.99.102:8031/WXJXJY/Pages/mycourses/MyCourses.aspx"
)

func main() {
	pw, browser, err := playwrightKit.LaunchBrowser(playwrightKit.BrowserNameFirefox, "_playwright-deps", true, &playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		panic(err)
	}
	defer pw.Stop()
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		panic(err)
	}
	if _, err = page.Goto(url); err != nil {
		logrus.Fatalf("could not goto: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		logrus.Infof("input: [%s]", text)

		switch text {
		case "quit":
			// 此时不会继续接受输入
			logrus.Info("you needn't to input text")
		case "pages":
			contexts := browser.Contexts()
			var pages []playwright.Page
			for _, context := range contexts {
				pages = append(pages, context.Pages()...)
			}
			logrus.Infof("len(pages): %d", len(pages))
			for _, page := range pages {
				logrus.Infof("page.URL(): %s", page.URL())
			}
		case "start":

			//log.Printf("Total number of pages: %d", len(pages))
			//for _, page := range pages {
			//	log.Printf("Page URL: %s", page.URL())
			//}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	logrus.Info("=== end ===")
}
