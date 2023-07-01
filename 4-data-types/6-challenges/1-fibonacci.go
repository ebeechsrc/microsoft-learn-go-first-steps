package main

import (
	"fmt"
	"strconv"
)

func generateFibonacciSlice(n int) []int {
	f := make([]int, n)
	f[0], f[1] = 0, 1
	for i := 2; i < n; i++ {
		f[i] = f[i-1] + f[i-2]
		// check for overflow
		if f[i] < f[i-1] {
			panic("Overflow happened at element: " + strconv.Itoa(i+1) + " of the Fibonacci slice, use a smaller number")
		}
	}
	return f
}

func getValidFibonacciRangeFromUser() int {
	var n int
	fmt.Scanf("%d", &n)
	if n < 2 {
		fmt.Println("The number of elements must be greater than 1")
		return getValidFibonacciRangeFromUser()
	}
	return n
}

func main() {
	fmt.Println("Generating Fibonacci slice")
	fmt.Println("Enter the number of elements: ")
	var n int
	n = getValidFibonacciRangeFromUser()
	f := generateFibonacciSlice(n)
	fmt.Println(f)
}
