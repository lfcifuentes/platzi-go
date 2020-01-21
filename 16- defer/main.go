package main

import "fmt"

// el defer ejecuta la funcion que se le pasa por parametro al finalizar la funcion que lo llamo
// el ultino defer declarado es el que se ejecuta
func main() {
	defer fmt.Println(test())
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1") // primero se ejecuta este defer por ser el ultimo
	fmt.Println("hello")
}

func test() string {
	return "world"
}
