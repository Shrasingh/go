//A channel in Go is a typed conduit (pipe) that allows goroutines to communicate and share data safely.
//Why we use it:

//To send data between goroutines without using shared memory or mutexes.

//It synchronizes goroutines automatically: sending blocks until receiving and vice

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go")
	mych := make(chan int, 2) //buffered channel of capacity 2
	wg := &sync.WaitGroup{}

	mych <- 1
	mych <- 2
	// mych <- 3 // This would block since the channel is full
	fmt.Println(<-mych)
	fmt.Println(<-mych)
	// fmt.Println(<-mych) // This would block since the channel is empty
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		mych <- 3
		mych <- 4
		close(mych)
		fmt.Println(<-mych)
		fmt.Println(<-mych)
	}(mych, wg)
	wg.Add(1)
	wg.Wait()
	fmt.Println(<-mych)
	close(mych) // close the channel when done
}
