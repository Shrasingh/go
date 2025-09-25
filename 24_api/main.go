package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"math/rand"

	"github.com/gorilla/mux"

	"strconv"
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
	r := mux.NewRouter()

	// routes - file, routes.go
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	// seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", Price: 299, Author: &Author{Fullname: "John Doe", Website: "johndoe.com"}, Tags: []string{"frontend", "javascript"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "MERN Stack", Price: 199, Author: &Author{Fullname: "Jane Doe", Website: "janedoe.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Angular", Price: 299, Author: &Author{Fullname: "John Smith", Website: "johnsmith.com"}, Tags: []string{"frontend", "typescript"}})
	//listen to a port

	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")
	//r.HandleFunc("/courses/search", searchCourse).Methods("GET")
	// starting server
	fmt.Printf("Starting server at port 4000\n")
	if err := http.ListenAndServe(":4000", r); err != nil {
		panic(err)
	}
	// http://localhost:4000

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
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	// generate unique id, string
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) // mock id - not safe
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	// loop through courses, find matching id
	// if no id, send response: "No course found with given id"
	// else send the response
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			if course.IsEmpty() {
				json.NewEncoder(w).Encode("Please send some data")
				return
			}
		}
	}
	// if no id, send response: "No course found with given id"
	w.Write([]byte("No course found with given id"))
	json.NewEncoder(w).Encode("No course found with given id")
	return

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	// loop through courses, find matching id
	// if no id, send response: "No course found with given id"
	// else send the response
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(courses)

			return
		}
	}
	// if no id, send response: "No course found with given id"
	w.Write([]byte("No course found with given id"))
	json.NewEncoder(w).Encode("No course found with given id")
	return
}
func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all courses")
	w.Header().Set("Content-Type", "application/json")
	courses = nil
	json.NewEncoder(w).Encode(courses)
}
