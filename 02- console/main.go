package main

import "fmt"

func main() {
	var name string
	var lastname string
	fmt.Print("Ingresa tu nombre:")
	fmt.Scanf("%s", &name)

	fmt.Print("Ingresa tu apellido:")
	fmt.Scanf("%s", &lastname)

	fmt.Printf("Hello %s %s", name, lastname)
}
