package name

import "fmt"

// si la primera letra de la función es minuscula, la función es privada
func test() {
	fmt.Println("Función privada")
}

// GetName optiene y retorna el nombre de una persona
func GetName() string {
	var name string

	fmt.Print("Ingresa tu nombre:")
	fmt.Scanf("%s", &name)

	return name
}
