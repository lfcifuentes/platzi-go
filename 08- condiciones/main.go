package main

import "fmt"

func main() {
	var number int
	fmt.Println("Ingrese un numero:")
	fmt.Scanf("%d", &number)

	if number == 0 {
		fmt.Println("El numero es cero")
	} else if number%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}

}
