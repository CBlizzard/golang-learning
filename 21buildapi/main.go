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

func main() {
	fmt.Println("rolling thundaaaaaaa")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseID: "1", CourseName: "way of the ace", CoursePrice: 500, Author: &Author{Fullname: "Bokuto Kotaro", Website: "bokutoace.com"}})
	courses = append(courses, Course{CourseID: "2", CourseName: "serve to win", CoursePrice: 1000, Author: &Author{Fullname: "Oikawa Toru", Website: "fakepride.com"}})

	// routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Model

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

// fake DB
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

// Controller

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hey hey hey</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	// setting header
	w.Header().Set("Content-Type", "application/json")

	// setting response
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	// setting header
	w.Header().Set("Content-Type", "application/json")

	// get params
	params := mux.Vars(r)
	fmt.Println(params) // delete this after see the result

	// get required course
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// if not found
	json.NewEncoder(w).Encode("Course not found")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	// setting header
	w.Header().Set("Content-Type", "application/json")

	// no data
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data in body")
		return
	}

	// empty data ie {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Empty data")
		return
	}

	// generating ID
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	// setting header
	w.Header().Set("Content-Type", "application/json")

	// get ID
	params := mux.Vars(r)

	// since no real DB is used, ungli tedhi karni padegi
	// loop & get index, remove old value, then add updated value
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseID = params["id"]
			courses = append(courses, updatedCourse)

			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	// setting header
	w.Header().Set("Content-Type", "application/json")

	// get ID
	params := mux.Vars(r)

	// loop & get index, then delete
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted")
			return
		}
	}
}
