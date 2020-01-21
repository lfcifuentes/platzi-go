package main

import "fmt"

// Platzi empresa con cursos
type Platzi struct {
	Name    string
	Courses []PlatziCourse
}

// PlatziCourse estructura para cursos con platzi
type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

// GetCourses funcion que imprime los cursos de platzi
func (p *Platzi) GetCourses() {
	for _, detail := range p.Courses {
		fmt.Println(detail.Name)
	}
}

func main() {

	platzic1 := PlatziCourse{
		Name:   "Go",
		Slug:   "go",
		Skills: []string{"1", "2", "3"},
	}

	fmt.Println(platzic1)

	platzic2 := new(PlatziCourse)
	platzic2.Name = "Gitlab"
	platzic2.Slug = "gitlab"
	platzic2.Skills = []string{"DevOPS"}

	fmt.Println(platzic2)

	platzic3 := PlatziCourse{
		Name:   "php",
		Slug:   "php",
		Skills: []string{"1", "2", "3"},
	}

	platzi := new(Platzi)
	platzi.Name = "Platzi"
	platzi.Courses = []PlatziCourse{platzic1, platzic3}

	fmt.Println(platzi)

	platzi.GetCourses()
}
