package main

import (
	"fmt"
	"github.com/samber/lo"
	"math"
)

func main() {
	fmt.Println(lo.Elipse("Richelieu", 0))
	fmt.Println(lo.Elipse("Richelieu", 1))
	fmt.Println(lo.Elipse("Richelieu", 2))
	fmt.Println(lo.Elipse("Richelieu", 3))
	fmt.Println(lo.Elipse("Richelieu", 5))
	fmt.Println(lo.Elipse("Richelieu", math.MaxInt))
}
