package main

import (
	"fmt"
)

func main() {
	fmt.Println("IF ELSE in Golang \n")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Watch out"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println(" Nums is less than 10 ")
	} else {
		fmt.Println(" Not less than 10 ")
	}

	// if err != nil {

	// }


}
