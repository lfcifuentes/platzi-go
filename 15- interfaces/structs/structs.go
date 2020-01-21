package structs

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

// Measure interfase de geometria
func Measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

//Rect estructura rectanglo
type Rect struct {
	Width  float64
	Height float64
}

// Circle estructura circulo
type Circle struct {
	Radius float64
}

func (r Rect) area() float64 {
	return r.Width * r.Height
}
func (r Rect) perim() float64 {
	return 2*r.Width + 2*r.Height
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.Radius
}
