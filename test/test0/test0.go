package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

var (
	port = 9000
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, strconv.Itoa(port))
	})

	if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		logrus.Fatal(err)
	}
}
