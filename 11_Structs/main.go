package main

import "fmt"

func main() {
	fmt.Println("Structs in Go lang")

	// No inheritance in Golang; No super / parent

	vineet := User{"Vineet", "vineet@go.com", true, 21}

	// fmt.Println("Struct: ", vineet, "\n")
	// fmt.Printf("Vineet details are : %+v", vineet)

	fmt.Printf("Name: %v \n", vineet.Name)
	fmt.Printf("Email: %v \n", vineet.Email)
	fmt.Printf("Age: %v \n", vineet.Age)
	fmt.Printf("Status: %v \n", vineet.Status)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
