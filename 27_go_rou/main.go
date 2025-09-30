package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Start goroutines
	go greater("hellow")
	go greater("shraddha")

	// Prevent main from exiting immediately
	// Option 1: Just wait long enough
	time.Sleep(20 * time.Second)

	// Option 2 (better): use sync.WaitGroup (more controlled)
	// Uncomment below if you want the clean solution
	/*
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			greater("hellow")
		}()

		go func() {
			defer wg.Done()
			greater("shraddha")
		}()

		wg.Wait()
	*/
	websitelist := []string{
		"https://www.google.com",
		"https://www.shraddha.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
	}
	var wg sync.WaitGroup // usually they are pointers but here it's fine
	// Add the number of goroutines to wait for
	wg.Add(len(websitelist))

	for _, web := range websitelist {
		go func(url string) {
			defer wg.Done()
			getStatusCode(url)
		}(web)
	}

	wg.Wait() // wait for all goroutines to finish
}

func greater(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops, error in endpoint:", endpoint, "-", err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("Status code for %s: %d\n", endpoint, res.StatusCode)
}
