package main

import (
	"fmt"

	geometry "example.com/geometry"
)

func printInformation(s geometry.Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

func main() {
	var s geometry.Shape = geometry.NewSquare(3)
	printInformation(s)

	c := geometry.NewCircle(6)
	printInformation(c)
}
