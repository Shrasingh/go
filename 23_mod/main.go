package main // main package - entry point of the Go program

import (
	"fmt"      // for printing messages to console
	"log"      // for logging errors
	"net/http" // for creating HTTP server

	"github.com/gorilla/mux" // third-party router for handling routes
)

func main() {
	fmt.Println("Welcome to mod") // prints a message on terminal
	Greater()                     // calls the Greater() function

	r := mux.NewRouter() // creates a new router instance from Gorilla Mux

	// register a route: if user visits "/", call serveHome function, but only for GET requests
	r.HandleFunc("/", serveHome).Methods("GET")

	// start the server on port 4000 with router 'r'
	// log.Fatal ensures if server crashes, the error is logged and program exits
	log.Fatal(http.ListenAndServe(":4000", r))
}

func Greater() {
	fmt.Println("hey there mod users") // simple function printing a message
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// sends an HTML response to the browser
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}
