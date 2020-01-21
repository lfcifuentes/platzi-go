package main

import "fmt"

func main() {
	// si un if tiene mas de 2 condiciones es mas recomendable cambiarlo a un switch
	var number = 0
	fmt.Println("Ingrese un número del 1 al 10:")
	fmt.Scanf("%d", &number)
	// con valor
	switch number {
	case 1:
		fmt.Println("El numero es 1")
	default:
		fmt.Println("El numero no es 1")
	}
	// con condicionales
	switch {
	case number%2 == 0:
		fmt.Println("El numero es par")
	default:
		fmt.Println("El numero es impar")
	}

}
