package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"
	"log"
	"net/http"
	"strings"
)

func main() {
	g := gin.Default()

	gzipHandler := gzip.NewHandler(gzip.Config{
		CompressionLevel: 1,
		MinContentLength: 1 * 1024,
		RequestFilter: []gzip.RequestFilter{
			gzip.NewCommonRequestFilter(),
			//gzip.DefaultExtensionFilter(),
		},
		ResponseHeaderFilter: []gzip.ResponseHeaderFilter{
			gzip.NewSkipCompressedFilter(),
			gzip.DefaultContentTypeFilter(),
		},
	})
	g.Use(gzipHandler.Gin)

	g.GET("/api.do", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"msg":  "hello",
			"data": fmt.Sprintf("l%sng!", strings.Repeat("o", 100)),
		})
	})

	log.Println(g.Run(fmt.Sprintf(":%d", 3000)))
}
