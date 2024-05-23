package main

import (
	"fmt"
	"github.com/itchyny/gojq"
)

func main() {
	query, err := gojq.Parse(".name")
	if err != nil {
		panic(err)
	}
	input := map[string]interface{}{
		"name": "Alice",
		"age":  30,
		"city": "Wonderland",
	}
	iter := query.Run(input)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(v)
	}
	/*
		output: Alice
	*/
}
