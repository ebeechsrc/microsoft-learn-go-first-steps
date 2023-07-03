package main

import (
	"fmt"

	store "example.com/store"
)

func main() {
	var employeeAccount store.Employee
	employeeAccount.ChangeName(store.Name{"John", "Doe"})
	fmt.Println(employeeAccount.Name())
	printCredits(employeeAccount)
	employeeAccount.AddCredits(10)
	printCredits(employeeAccount)
	employeeAccount.RemoveCredits(5)
	printCredits(employeeAccount)
}

func printCredits(employeeAccount store.Employee) {
	fmt.Printf("Credits: %v\n", employeeAccount.CheckCredits())
}
