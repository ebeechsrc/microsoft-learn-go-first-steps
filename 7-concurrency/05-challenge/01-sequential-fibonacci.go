package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(results chan<- float64, number float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	results <- x
}

func main() {
	start := time.Now()
	fibonacci_results := make(chan float64, 10)

	size := 15
	for i := 1; i < size; i++ {
		go fib(fibonacci_results, float64(i))
	}
	for i := 1; i < size; i++ {
		n := <-fibonacci_results
		fmt.Printf("%v\n", n)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
