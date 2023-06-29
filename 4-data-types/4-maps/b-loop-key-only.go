package main

import (
	"fmt"
)

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	for name := range studentsAge {
		fmt.Printf("Names %s\n", name)
	}
	john, johnIsAlive := studentsAge["john"]
	if johnIsAlive {
		fmt.Printf("John's age is %d\n", john)
	}
}
