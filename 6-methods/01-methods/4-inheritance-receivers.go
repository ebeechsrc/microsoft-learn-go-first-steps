package main

import "fmt"

type triangle struct {
	size int
}

type coloredTriangle struct {
	triangle
	color string
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

// this isn't needed as the coloredTriangle struct inherits the perimeter method
// this is a sill example...
// func (t coloredTriangle) perimeter() int {
// 	return t.triangle.perimeter()
// }

func main() {
	t := coloredTriangle{triangle{3}, "blue"}
	t.triangle.doubleSize()
	t.doubleSize()
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter", t.perimeter())
}
