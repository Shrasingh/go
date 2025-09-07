package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch and case in golang")
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("dice number is ", dicenumber)
	switch dicenumber {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
		fallthrough //In Go, fallthrough forces execution to continue into the next case block regardless of its condition.
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
	default:
		fmt.Println("this is default case")
	}

}
