package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("welcome to files in golang")
	content := "this needs to go in a file - shraddha"
	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is ", length)
	defer file.Close()
	readFile(("./myfile.txt"))

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("text in file is \n", databyte)
	fmt.Println("text in file is \n", string(databyte))
	fmt.Println("text in file is \n", strings.Split(string(databyte), "\n"))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
