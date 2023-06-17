package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

// Model for courses - file

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

// fake db
var courses []Course

// middleware helper-file
func (c *Course) IsEmpty() bool { //returns a true value if both are empty
	return c.CourseName == ""
}

type Author struct {
	Fullname string `json:"fullname"`
	Websiite string `json:"website"`
}

func main() {
	// use slice for database
	fmt.Println("API is running")

	// create a router
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Go", CoursePrice: 100, Author: &Author{Fullname: "John", Websiite: "john.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Python", CoursePrice: 200, Author: &Author{Fullname: "Doe", Websiite: "doe.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Java", CoursePrice: 300, Author: &Author{Fullname: "Smith", Websiite: "smith.com"}})

	// create a route
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
	
}

// conroller - file this means that controller goes in its seperate file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}

// through up all teh data from the database
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	//  to add header
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	// grab that id
	// comapre the value using loop
	// if matches return the value
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r) // user is sending the id and is stored in the params
	fmt.Printf("Type of Params is: %T\n", params)

	// loop through and find the id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with that id ")
	return
}

// Adding the course in database

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add one course")
	w.Header().Set("Content-Type", "application/json")

	// what if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about -{} empty object
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	fmt.Printf("Type of course is: %T\n", course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send the data in correct format")
		return
	}

	// generate unique id, convert it to the string
	// append course into the string

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// update a course from the database
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update the course")
	w.Header().Set("Content-Type", "application/json")

	// first grab the id from the request
	params := mux.Vars(r)
	fmt.Printf("Type of params is: %T\n", params)

	//loop, get teh id, and remove the existing id and addd with the new id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO : Send the error message
}

// delete a course from the database
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop through and find the id, remove it, (index , index+1) take care of the index
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			fmt.Println("Course deleted successfully")
			break
		}
	}
	json.NewEncoder(w).Encode(courses)
	return

}
