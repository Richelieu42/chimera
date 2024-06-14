package playwrightKit

import (
	"github.com/playwright-community/playwright-go"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"os"
	"strings"
)

// LaunchBrowser
/*
@param browserName 		浏览器名称
@param tempDirPath 		用于存放浏览器和操作系统的依赖
@param launchOptions 	可以为nil，将采用默认值(headless)
*/
func LaunchBrowser(browserName string, tempDirPath string, installFlag bool,
	launchOptions *playwright.BrowserTypeLaunchOptions) (pw *playwright.Playwright, browser playwright.Browser, err error) {

	if err = fileKit.AssertNotExistOrIsDir(tempDirPath); err != nil {
		return
	}

	if browserName == "" {
		browserName = BrowserNameChromium
	}
	browserName = strings.ToLower(browserName)
	switch browserName {
	case BrowserNameChromium, BrowserNameFirefox, BrowserNameWebkit:
	default:
		err = errorKit.Newf("invalid browserName(%s)", browserName)
		return
	}

	runOptions := &playwright.RunOptions{
		DriverDirectory:     tempDirPath,
		SkipInstallBrowsers: false,
		Browsers:            []string{browserName},
		Verbose:             true,
		Stdout:              os.Stdout,
		Stderr:              os.Stderr,
	}

	if installFlag {
		err = playwright.Install(runOptions)
		if err != nil {
			return
		}
	}
	pw, err = playwright.Run(runOptions)
	if err != nil {
		return
	}

	var tmp []playwright.BrowserTypeLaunchOptions = nil
	if launchOptions != nil {
		tmp = append(tmp, *launchOptions)
	}
	switch browserName {
	case BrowserNameChromium:
		browser, err = pw.Chromium.Launch(tmp...)
	case BrowserNameFirefox:
		browser, err = pw.Firefox.Launch(tmp...)
	case BrowserNameWebkit:
		browser, err = pw.WebKit.Launch(tmp...)
	default:
		// Richelieu: 理论上代码不会走到此处
		err = errorKit.Newf("invalid browserName(%s)", browserName)
	}
	return
}
