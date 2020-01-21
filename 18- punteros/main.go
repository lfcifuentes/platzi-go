package main

import "fmt"

func main() {
	a := 100
	var b *int

	b = &a

	fmt.Println(a)
	fmt.Println(b)

	a++
	*b++

	fmt.Println(a)
	fmt.Println(*b)
}
