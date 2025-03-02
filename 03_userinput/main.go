package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// welcome := "Welcome to Go user input"
	// fmt.Println(welcome)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter the rating for our Pizza : ")

	// // comma ok syntax or error ok syntax

	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thanks for rating, ", input)

	// fmt.Println("Enter the rating for our Pizza : ")

	// // input, _ := reader.ReadString('\n')

	// _, err := strconv.Atoi(input)
	// fmt.Printf("Type of rating is: %T", inputType)

	// fmt.Println("Error is : ", err)
	// fmt.Printf("Type of err is : %T", err)

	// myMap := map[string]string{
	// 	"name": "John Doe",
	// 	"Age": "33",
	// 	"Gender": "Mechanic",
	// }

	// name, ok := myMap["nasme"]
	// fmt.Printf("Type of ok is : %T",ok)
	// fmt.Println(name)

	// VOTING SYSTEM

	var greetings string = "Welcome to the Voting System"
	fmt.Println(greetings)

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the age for voting : ")
	age, _ := inputReader.ReadString('\n')
	age = strings.TrimSpace(age)

	numericAge, err := strconv.Atoi(age)

	if err == nil {
		fmt.Printf("Your age is %d \n", numericAge)

		if numericAge < 18 {
			fmt.Println("Your are not eligible to vote")
		} else {
			fmt.Println("You are eligible to vote")
			fmt.Println("!!!! Thanks for Voting !!!")
		}

	} else {
		fmt.Println("Error is : ", err)
		fmt.Println("Please enter a valid number")
	}

}
