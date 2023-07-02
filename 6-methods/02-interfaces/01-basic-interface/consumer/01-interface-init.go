package main

import (
	"fmt"

	geometry "example.com/geometry"
)

func main() {
	var s geometry.Shape = geometry.NewSquare(3)
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
