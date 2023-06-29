package main

import "fmt"

func main() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]
	slice2 := letters[1:4]

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)
}
