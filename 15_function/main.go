package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()
	greeterTwo()
	result := adder(3, 5)
	fmt.Println("the result is ", result)
	result = adder(7, 9)
	fmt.Println("the result is ", result)
	proRes := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("the pro result is ", proRes)
}
func greeter() {
	fmt.Println("hello world from greeter function")
}
func greeterTwo() {
	fmt.Println("hello world from greeterTwo function")
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
