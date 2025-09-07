package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")
	var ptr *int
	fmt.Println("the value of pointer is ", ptr)
	myNumber := 23
	var ptr1 = &myNumber
	fmt.Println("the value of pointer is ", ptr1)
	fmt.Println("the value of pointer is ", *ptr1)
}
