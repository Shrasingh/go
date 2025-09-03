package main

import "fmt"

const LoginToken string = "this is a login token" // public
// const are globaly scoped
func main() {
	var username string = "shraddha"
	fmt.Println(username)
	//fmt.Printf(username)
	fmt.Printf("Variable is of type: %T \n", username)
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
	var smallFloat float32 = 255.4555555555
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)
	var anotherFloat float64
	fmt.Println(anotherFloat)
	fmt.Printf("Variable is of type: %T \n", anotherFloat)

	// implitit typeing
	var website = "learncodeonline.in"
	fmt.Println(website)
	// no var style
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)
	// outside of a method := is not allowed
	fmt.Printf("Login token is: %s \n", LoginToken)
}
