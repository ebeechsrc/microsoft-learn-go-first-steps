package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee)
	employee.FirstName = "James"
	fmt.Println(employee)
	// not a copy, it's a reference... bad names...
	employeeCopy := &employee
	employeeCopy.FirstName = "David"
	fmt.Println(employee)
}
