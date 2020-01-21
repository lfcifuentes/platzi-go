package main

import (
	"fmt"

	"./calculator"
)

func main() {
	response, err := calculator.Sum(1, "2")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
