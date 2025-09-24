package main

import "fmt"

func main() {
	defer fmt.Println("welcome to defer in golang")
	fmt.Println("welcome to golang")
	// multiple defer statements are executed in LIFO order
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("learning defer")
	// defer fmt.Println("this is the end")
	// panic("something went wrong")
	fmt.Println("end of the main function")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
