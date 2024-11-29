package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"log"
	"os"
)

func init() {
	log.Printf("pid: %d", os.Getpid())
}

func init() {
	for {
	}
}

func main() {
	log.Println("--- main starts ---")

	port := 8002

	//console.PrintBasicDetails()

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}

	log.Println("--- main ends ---")
}
