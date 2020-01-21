package main

import "fmt"

func main() {
	mapa1 := getMap()

	fmt.Println(mapa1["llave1"]) // retorna el valor del mapa

	fmt.Println(mapa1["llave3"]) // si la llave no existe retorna un valor por defecto 0 si es entero, "" si es string etc

	getMapTwo()

}

// getMap devuelve un map con llave valor
func getMap() map[string]string {
	mapTest := make(map[string]string) // asignar el valor a la variable e inicializarlo
	mapTest["llave1"] = "18"
	mapTest["llave2"] = "323324123"
	return mapTest
}

// getMap devuelve un map con llave valor
func getMapTwo() map[string]int {
	mapTest := map[string]int{
		"luis":  23,
		"pedro": 19,
		"maria": 29,
	}

	fmt.Println(mapTest)

	delete(mapTest, "luis")

	fmt.Println(mapTest)

	return mapTest
}
