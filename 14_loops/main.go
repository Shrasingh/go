package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on loops")
	days := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	fmt.Println("days are ", days)
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	for i, day := range days {
		fmt.Printf("index is %v and day is %v\n", i, day)
	}
	for _, day := range days {
		fmt.Printf("day is %v\n", day)
	}
	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			break
		}
		fmt.Println("rougue value is ", rougueValue)
		rougueValue++
	}
lco:
	fmt.Println("Jumping at")
}
