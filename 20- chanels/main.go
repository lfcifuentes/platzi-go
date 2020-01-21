package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		defer close(ch)
		ch <- "Hola chanel" // se esta enviando al canal
	}()

	fmt.Println(<-ch) // significa que el canal esta enviando la información

	ints := make(chan int)

	go func() {
		defer close(ints)
		ints <- 1
		ints <- 2
		ints <- 3
		ints <- 4
		ints <- 5
	}()

	for i := range ints {
		fmt.Println(i)
	}

	twoInts := make(chan int, 2) // definimos que el canal es de maximo dos datos

	twoInts <- 1
	twoInts <- 2

	fmt.Println(<-twoInts)
	fmt.Println(<-twoInts)
	// para enviarle mas informacion al canal debo eliminar almenos un dato del buffer con <-buffer
	twoInts <- 3
	fmt.Println(<-twoInts)

}
