package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Welcome to user input program"
	fmt.Println(welcomeMessage)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for pizza:")
	// comma ok // err ok
	// ip err
	input, _ := reader.ReadString('\n')
	// input,err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T \n", input)
}

// bufio
