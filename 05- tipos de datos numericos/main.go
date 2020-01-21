package main

import "fmt"

func main() {
	var intOne int     // maximo de 32 bits   -> int32
	var intTwo uint    // numero positivo  -> int64
	var intThree int64 // numero de 64 bits

	fmt.Println(intOne)
	fmt.Println(intTwo)
	fmt.Println(intThree)

	var float32 float32 // flotante de 32 bits
	var float64 float64 // flotante de 64 bits

	fmt.Println(float32)
	fmt.Println(float64)
}
