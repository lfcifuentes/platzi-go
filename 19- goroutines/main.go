package main

import (
	"fmt"
	"time"
)

func main() {

	go forRoutines("uno", 500)
	go forRoutines("dos", 400)

	time.Sleep(10000 * time.Millisecond)
}

func forRoutines(name string, index int) {
	for i := 0; i < index; i++ {
		go routine(name, i)
	}
}
func routine(name string, index int) {
	fmt.Printf("ESte es el hilo numero %d de %s \n", index, name)
}
