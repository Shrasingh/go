package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://shraddhasingh.vercel.app/"

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("lco web request")
	fmt.Println(myurl)
	// parsing the url
	// we use net/url package to parse url
	// import "net/url"
	//
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Fragment)
	fmt.Println(result.RawPath)

	qparams := result.Query()
	fmt.Printf("type of query params is : %T\n", qparams)
	fmt.Println(qparams["homepage"])
	for _, val := range qparams {
		fmt.Println("param is ", val)
	}
	// building url from scratch
	// & is used to append multiple query params
	// = is used to assign value to a key
	// ? is used to start query params
	//
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println("another url is ", anotherUrl)

}
