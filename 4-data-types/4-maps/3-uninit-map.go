package main

import "fmt"

func main() {
	var studentsAge map[string]int
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println(studentsAge)
}
