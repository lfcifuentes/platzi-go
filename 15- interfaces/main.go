package main

import (
	"./structs"
)

func main() {
	r := structs.Rect{Width: 3, Height: 4}
	structs.Measure(r)
	c := structs.Circle{Radius: 5}
	structs.Measure(c)
}
