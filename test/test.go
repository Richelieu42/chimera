package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
