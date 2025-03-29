package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"math/rand"

	"github.com/gorilla/mux"
)

type course struct {
	CourseID    string  `json:courseID`
	CourseName  string  `json:courseName`
	CoursePrice float64 `json:coursePrice`
	Author      *Author `json:author`
}

type Author struct {
	Name    string `json:name`
	Website string `json:website`
}

// slice of structs as fake database
var courses []course

// middleware,helper functions
func (c *course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == "" && c.CoursePrice == 0 && c.Author == nil
}

//controller functions- file (different functions for different routes)

// serveHome route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welocome to API by HRK</h1>"))
}

//get all courses

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome to get all the courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

// GET ONE COURSE BASES ON REQUEST ID
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	//grab all the variable  from client request (id etc)  and store it in params which is of key-val pair

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given id")

}

// Add a course to controller
func addOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcom to a new course adder")
	w.Header().Set("content-type", "application/json")

	//what if body is empty (no content added)
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data inside")
	}

	var course course

	json.NewDecoder(r.Body).Decode(&course)

	//generate unique id , convert into type string , append the course into courses slice

	//rand.Intn(1000) generates a random number between 1 to 1000
	//rand.Seed(time.Now().UnixNano())
	//course.CourseID = strconv.Itoa(rand.Intn(1000))

	rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseID = strconv.Itoa(rand.Intn(1000))

	courses = append(courses, course)

	//send this back in the form of json

	json.NewEncoder(w).Encode(course)

}

// update a course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating a course")
	w.Header().Set("content-type", "application/json")

	//first - grab id form request , remove that course  , add new course with new id
	params := mux.Vars(r)

	if params["id"] == "" {
		json.NewEncoder(w).Encode("No id sent, pls send with a id")
	}

	for index, mycourse := range courses {
		if mycourse.CourseID == params["id"] {
			//removed that particular course
			courses = append(courses[:index], courses[index+1:]...)
			break

		}

	}
	var course course
	json.NewDecoder(r.Body).Decode(&course)

	course.CourseID = params["id"]
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// delete a course
func deletecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting a course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for idx, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			break
		}
	}

	message := fmt.Sprintf("successfully deleted a course having id: %s", params["id"])
	json.NewEncoder(w).Encode(message)

}

func main() {
	fmt.Println("welcome to func main")

	//Handling routes and testing routes
	r := mux.NewRouter()

	//seeding (testing with somedata)
	courses = append(courses, course{CourseID: "1", CourseName: "ReactJs by Salman", CoursePrice: 9, Author: &Author{Name: "Salman Khan", Website: "learnwithSalman.com"}})
	courses = append(courses, course{CourseID: "2", CourseName: "python by Shahrukh", CoursePrice: 10, Author: &Author{Name: "Shahrukh khan", Website: "learnwithShahrukh.com"}})
	courses = append(courses, course{CourseID: "3", CourseName: "javascript with Amir", CoursePrice: 11, Author: &Author{Name: "AmirKhan", Website: "learnwithAmirKhan.com"}})
	//courses = append(courses , course{CourseID: "4" , CourseName: "" , CoursePrice: 9 , Author: &Author{Name: "Salman Khan" , Website: "learnwithSalman.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/saarecourse", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	//use POST when you want to create a new resource and PUT when you want to update or replace an existing resource.
	r.HandleFunc("/course", addOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deletecourse).Methods("DELETE")

	//listener to a port

	log.Fatal(http.ListenAndServe(":2000", r))
}
