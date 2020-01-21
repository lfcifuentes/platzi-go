package main

import "fmt"

func main() {
	// for normal
	for i := 0; i < 5; i++ {
		fmt.Printf("I es: %d \n", i)
	}
	// for con condición
	c := 100
	for c > 0 {
		c -= 10
		fmt.Printf("Otro for: %d \n", c)
	}
	// form infinito
	s := 1000
	for {
		s -= 1
		if s == 0 {
			fmt.Println("El for infinito termino")
			break
		}
	}
}
