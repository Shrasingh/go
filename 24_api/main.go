package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId   string   `json:"courseid"`
	CourseName string   `json:"coursename"`
	Price      int      `json:"price"`
	Author     *Author  `json:"author"`
	Tags       []string `json:"tags,omitempty"`
}
type Author struct {
	Fullname string
	Website  string
}

// fake db
var courses []Course

// middleware helper - file, main.go
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseId == "" && c.CourseName == ""
}
func main() {
	fmt.Println("Welcome to API video")
}

// controllers - file, controllers.go
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	// loop through courses, find matching id
	// if no id, send response: "No course found with given id"
	// else send the response
	// grab id from request
	params := mux.Vars(r)
	// loop through courses, find matching id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// if no id, send response: "No course found with given id"
	w.Write([]byte("No course found with given id"))
	json.NewEncoder(w).Encode("No course found with given id")
	return
}
