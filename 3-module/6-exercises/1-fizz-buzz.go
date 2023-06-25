package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(value int) string {
	// fizz buzz example
	isMatch := func(denominator int) bool {
		return value%denominator == 0
	}

	divisibleBy3 := isMatch(3)
	divisibleBy5 := isMatch(5)

	switch {
	case divisibleBy3 && divisibleBy5:
		return "FizzBuzz"
	case divisibleBy3:
		return "Fizz"
	case divisibleBy5:
		return "Buzz"
	}
	return strconv.Itoa(value)
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
