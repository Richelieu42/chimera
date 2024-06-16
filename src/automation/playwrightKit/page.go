package playwrightKit

import (
	"github.com/playwright-community/playwright-go"
)

// GetPages
/*
@param browser 不能为nil
@return 有可能len(pages) == 0
*/
func GetPages(browser playwright.Browser) (pages []playwright.Page) {
	contexts := browser.Contexts()
	for _, context := range contexts {
		pages = append(pages, context.Pages()...)
	}
	return
}
