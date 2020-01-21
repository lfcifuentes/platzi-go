package main

import "fmt"

const helloWorld string = "Hello world %s"
const otherConst = "Other"

func main() {
	var name string = "test"
	fmt.Println(name)

	lastname := "Luis Fernando"
	fmt.Println(lastname)

	var (
		a = 1
		b = 2
		c = 3
	)

	fmt.Println(a, b, c)

	fmt.Printf(helloWorld, lastname)

	fmt.Printf("\n %s \n", otherConst)
}
