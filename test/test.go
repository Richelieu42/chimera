package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"

	"github.com/playwright-community/playwright-go"
)

func init() {
	//if err := playwright.Install(runOptions); err != nil {
	//	panic(err)
	//}
}

func main() {
	tempDirPath := "_temp"
	if err := fileKit.MkDirs(tempDirPath); err != nil {
		logrus.Fatal(err)
	}
	runOptions := &playwright.RunOptions{
		DriverDirectory:     tempDirPath,
		SkipInstallBrowsers: false,
		Browsers:            []string{"chromium"},
		Verbose:             true,
		Stdout:              os.Stdout,
		Stderr:              os.Stderr,
	}

	pw, err := playwright.Run(runOptions)
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://news.ycombinator.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	time.Sleep(time.Second * 10)

	page.Reload()

	entries, err := page.Locator(".athing").All()
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	for i, entry := range entries {
		title, err := entry.Locator("td.title > span > a").TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, title)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
