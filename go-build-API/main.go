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

//Model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	website  string `json:"website"`
}

// fake dataBase

var courses []Course

//helper - file(it mean it shoud in different file)
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to Golang API")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReatJS", CoursePrice: 299, Author: &Author{FullName: "Sourav", website: "souravsk.github.io"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Golang", CoursePrice: 199, Author: &Author{FullName: "Sourav", website: "souravsk.github.io"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":8080", r))
}

//controller - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome To API by Sourav</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through course find matching id and return the respone
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("NO Couse Found with Given id")
	return
}

// create one course

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("content-Type", "application/json")

	//what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Plese send some data")
	}

	//what about this {} this is data some kind of data but not a really data
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("NO data found inside JSON")
		return
	}

	//generate unique id, string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode("New Data is added into database")
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updataing one Value")
	w.Header().Set("content-Type", "application/json")

	//first grab id form request
	params := mux.Vars(r)

	//next loop and stop when you find id then remove that value and then add value into the same ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Data id Updated")
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO:
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one Course")
	w.Header().Set("content-Type", "application/json")

	//first grab id from request
	params := mux.Vars(r)

	//next loop and stop when you find id then remove that value and then add value into the same ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course id deleted")
			break
		}
	}
}
