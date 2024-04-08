package main

import (
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"html"
)

func main() {
	fmt.Println(html.UnescapeString("&lt;p&gt;Some text with &#39;quotes&#39; and &lt;em&gt;markup&lt;/em&gt;&lt;/p&gt;"))
}
