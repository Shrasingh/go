package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on switch")
	myFavNum := 7
	switch myFavNum {
	case 1:
		fmt.Println("1 is my fav number")
	case 7:
		fmt.Println("7 is my fav number")
	case 13:
		fmt.Println("13 is my fav number")
	default:
		fmt.Println("this is a defualt case")
	}

	fmt.Println("if else ")
	logincount := 10
	var result string
	if logincount < 10 {
		result = "less than 10"
	} else {
		result = "more than 10"
	}
	fmt.Println(result)
	if logincount < 10 {
		result = "less than 10"
	} else if logincount > 20 {
		result = "more than 20"
	} else {
		result = "more than 10 but less than 20"
	}

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}
	// imp
	if num := 3; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is more than 10")
	}
}
