package main

import "fmt"

func main() {

	var arrayOne [2]int

	arrayOne[0] = 1
	arrayOne[1] = 2

	fmt.Println(arrayOne)

	arrayTwo := [3]int{1, 2, 3}

	fmt.Println(arrayTwo)

	// slice
	var slice []int

	slice = append(slice, 4, 5)

	fmt.Println(slice)
}
