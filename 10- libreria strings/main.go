package main

import (
	"fmt"
	"strings"
)

func main() {
	var string = "Hello world, Hello platzi, Hello Go"

	// convertir a mayuscula
	upper := strings.ToUpper(string)
	fmt.Printf(upper + "\n")
	// convertir a minuscula
	lower := strings.ToLower(string)
	fmt.Println(lower)
	// remplazar textos
	stringReplace := strings.Replace(string, "Hello", "hola", -1)
	fmt.Println(stringReplace)
	// dividir una cadena
	stringSplit := strings.Split(string, " ")
	fmt.Println(stringSplit)

}
