package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(results chan<- int64, number int64) {
	var x, y, i int64 = 1, 1, 0

	for ; i < number; i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	results <- x
}

func get_input(quit_channel chan<- bool) {
	user_input := ""
	fmt.Scanf("%v", &user_input)
	should_i_keep_running := user_input != "quit"
	quit_channel <- should_i_keep_running
}

func main() {
	var exit_channel = make(chan bool, 5)
	var fibonacci_results = make(chan int64, 1)
	var fibonacci_nextIndexValue int64 = 0
	var fibonacci_lastIndexValue int64 = 0
	is_running := true
	start := time.Now()

	for is_running {
		go get_input(exit_channel)
		select {
		case n := <-fibonacci_results:
			fibonacci_lastIndexValue++
			fmt.Printf("%v\n", n)
		case is_running = <-exit_channel:
			if !is_running {
				continue
			}
			
			if fibonacci_nextIndexValue != fibonacci_lastIndexValue {
				fmt.Printf("I'm still working on the previous value %d, your input has been ignored, please wait.\n", fibonacci_nextIndexValue)
				continue
			}

			go fib(fibonacci_results, fibonacci_nextIndexValue)
			fibonacci_nextIndexValue++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
