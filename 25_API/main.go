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

type Course struct {
	Name     string  `json:"name"`
	Author   *Author `json:"author"`
	CourseID string  `json:"book_id"`
	Price    int     `json:"price"`
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

// !  Fake DB
var courses []Course

func (c *Course) checkEmpty() bool {

	return c.Name == ""
}

// ^ Controllers - newFile

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to API in Golang <h1>"))
}

func serveAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses function called: ")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// ! We extract params from the request (r) using mux.Vars(r)
	params := mux.Vars(r)

	// ^ Loop Through Courses
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}

	json.NewEncoder(w).Encode(" !!! Error: No book found with the given ID !!! ")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createOneCourse method called")
	w.Header().Set("Content-Type", "application/json")

	// @ 1) What if Body of request is empty ??
	if r.Body == nil {
		json.NewEncoder(w).Encode("!!!! Request Body Empty !!!!")
	}

	// What if {}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.checkEmpty() {
		json.NewEncoder(w).Encode("Struct is Empty ")
		return
	}

	for _, val := range courses {
		if val.Name == course.Name {
			json.NewEncoder(w).Encode("Course Already Exists")
			return
		}
	}

	// Generate RandomID -> convert it to string -> add it in Courses[] slice
	// rand.Seed(time.Now().UnixMicro()) // ! Depricated
	rnd := rand.New(rand.NewSource(time.Now().UnixMicro()))

	course.CourseID = strconv.Itoa(int(rnd.Int63n(1000)))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateOneCourse method called")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// todo: send a response when id is not found

	json.NewEncoder(w).Encode("ID NOT FOUND")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteOneCourse method called ! ! ! ")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Course Deleted Successfully !!! ")
			return
		}
	}

	// todo: if ID not found
	// ^ return errro
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("ERROR: ID NOT FOUND")

}

func main() {

	fmt.Println("Welcome to our Course API ")

	courses = append(courses, Course{
		CourseID: "2",
		Name:     "ReactJS",
		Price:    299,
		Author: &Author{
			Name:    "Vineet Salve",
			Website: "www.vineetsalve.com",
		},
	})

	courses = append(courses, Course{
		CourseID: "4",
		Name:     "MERN STack",
		Price:    199,
		Author: &Author{
			Name:    "John Cena",
			Website: "www.JC.com",
		},
	})

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", serveAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	fmt.Println("Server running ..... ")
	log.Fatal(http.ListenAndServe(":4000", r))

}
