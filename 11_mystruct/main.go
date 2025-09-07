package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on structs")
	shraddha := User{"shraddha", "shraddha@go.dev", 22, true}
	fmt.Println(shraddha)
	fmt.Printf("shraddha details are %+v\n", shraddha)
	fmt.Printf("name is %v and email is %v\n", shraddha.Name, shraddha.Email)
	shraddha.Age = 22
	fmt.Printf("updated age is %+v\n", shraddha)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

// there is no concept of classes in go, we use structs to create user defined types
