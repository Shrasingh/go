package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on arrays")
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "grapes"
	fruitList[3] = "peach"
	fmt.Println("fruit list is ", fruitList)
	fmt.Println("length of array is ", len(fruitList))
}
