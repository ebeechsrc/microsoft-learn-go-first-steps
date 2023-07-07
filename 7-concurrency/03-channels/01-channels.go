package main

import (
	fmt "fmt"
	http "net/http"
	"time"
)

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	ch <- fmt.Sprintf("CHECKING: %s is up and running!\n", api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func main() {
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet12345.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)
	start := time.Now()

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	// only the first message in the channel is printed.
	fmt.Print(<-ch)
	// time.Sleep(3 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
