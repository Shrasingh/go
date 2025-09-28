package main

import (
	"fmt"
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
}

func greater(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}
