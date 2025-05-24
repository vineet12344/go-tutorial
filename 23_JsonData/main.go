package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("--------WELCOME TO JSON TUTORIAL---------")
	fmt.Println("Course Details: ")
	encodeJson()
}

func encodeJson() {
	lcoCourses := []courses{
		{
			"Reactjs Bootcamp",
			399,
			"Udemy",
			"133456",
			[]string{},
		},
		{
			"Golang Bootcamp",
			299,
			"Udemy",
			"654654",
			[]string{"web-dev", "golang"},
		},
		{
			"Cloud Bootcamp",
			399,
			"Coursera",
			"1234",
			[]string{"cloud", "devops"},
		},
	}

	// Package this data as JSON Data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	checkError(err)

	fmt.Println("FINAL JSON: ", string(finalJson))

	// fmt.Printf("%s ", finalJson)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
