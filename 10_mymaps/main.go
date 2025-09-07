package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on maps")
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"
	fmt.Println("list of languages", languages)
	fmt.Println("js shorts for", languages["js"])
	delete(languages, "rb")
	fmt.Println("list of languages", languages)
	// loops
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
