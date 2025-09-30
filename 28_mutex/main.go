//A Mutex in Go (short for mutual exclusion) is a synchronization primitive used to protect shared resources from being accessed concurrently by multiple goroutines, preventing race conditions.
//In Go, multiple goroutines can run concurrently.

//If two or more goroutines try to read/write the same variable at the same time, it can cause inconsistent or incorrect data.

package main

import (
	"fmt"
	"sync"
)

// Shared counter variable
var counter int

// Mutex to protect the counter in concurrent access
var mu sync.Mutex

// incrementUnsafe increments the counter without using a mutex
// This is unsafe and may cause race conditions
func incrementUnsafe(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++ // unsafe: multiple goroutines can update counter simultaneously
	}
}

// incrementSafe increments the counter using a mutex
// This is safe for concurrent access
func incrementSafe(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock() // lock before accessing the shared resource
		counter++
		mu.Unlock() // unlock after done
	}
}

func main() {
	// --- Unsafe example ---
	counter = 0 // reset counter
	var wg1 sync.WaitGroup
	wg1.Add(2)

	go incrementUnsafe(&wg1)
	go incrementUnsafe(&wg1)

	wg1.Wait()
	fmt.Println("Counter without mutex (unsafe):", counter)

	// --- Safe example ---
	counter = 0 // reset counter
	var wg2 sync.WaitGroup
	wg2.Add(2)

	go incrementSafe(&wg2)
	go incrementSafe(&wg2)

	wg2.Wait()
	fmt.Println("Counter with mutex (safe):", counter)
}

//Without a mutex, counter may not reach 2000 because both goroutines may access and update it simultaneously.
//incrementUnsafe demonstrates race conditions: the counter may not reach 2000.

//incrementSafe uses mu.Lock() and mu.Unlock() to ensure mutual exclusion, so the counter is always correct.

//WaitGroup ensures main() waits for all goroutines to finish before printing.
