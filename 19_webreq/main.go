package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://shraddhasingh.vercel.app/"

func main() {
	fmt.Println("lco web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response is of type : %T\n", response)
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("data in bytes is ", databytes)
	fmt.Println("data in string is ", string(databytes))
	//
	fmt.Println("response is ", response)
	defer response.Body.Close()
	fmt.Println("response status is ", response.Status)
	fmt.Println("response status code is ", response.StatusCode)
	fmt.Println("response content length is ", response.ContentLength)
	fmt.Println("response headers are  ", response.Header)
}
