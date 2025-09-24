package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("web verbs in golang")
	PerformGetRequest()
	PerformPostRequest()
}

func PerformGetRequest() {
	fmt.Println("performing get request")
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code is ", response.StatusCode)
	fmt.Println("content length is ", response.ContentLength)
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, err := responseString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("byte count is ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostRequest() {
	fmt.Println("performing post request")
	const myurl = "http://localhost:8000/post"
	// fake json payload
	requestBody := strings.NewReader(`{
		"coursename":"lets go with golang",
		"price":0,
		"platform":"learncodeonline.in"
	}`)
	http.Post(myurl, "application/json", requestBody)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

// json payload
// {
// 	"coursename":"lets go with golang",

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("firstname", "shraddha")
	data.Add("lastname", "singh")
	data.Add("email", "	shraddha@go.dev")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
