package main

import (
	"fmt"

	"example.com/geometry"
)

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	// this will fail due to size being private
	// fmt.Println("Size", t.size)
	fmt.Println("Perimeter", t.Perimeter())
}
