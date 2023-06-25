package main

import "fmt"

func generateResponse(value int) string {
	switch {
	case value < 0:
		panic("You entered a negative number!")
	case value == 0:
		return ("0 is neither negative nor positive")
	default:
		return fmt.Sprintf("You entered: %v", value)
	}
}

func main() {
	for {
		fmt.Print("Enter number: ")
		value := 0
		fmt.Scanf("%d", &value)
		fmt.Println(generateResponse(value))
	}
}
