package main

import "fmt"

// Cannot use this outside of the function
// var is required to declare a variable
// var jwtToken := 3000000

// This is a package level variable
// This can be used in any function in this package
// This is a constant
// Since the first letter is capital this makes it public

const LoginToken string = "alkcmslkasmclksanacacois0ci3oaoia"

func main() {
	var username string = "John Doe"
	fmt.Println("Hello", username)
	fmt.Printf("Type of Varible is : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of Varible is : %T \n", isLoggedIn)

	var smallValue byte = 255
	fmt.Println("uint32", smallValue)
	fmt.Printf("Type of Varible is : %T \n", smallValue)

	var smallFloat float32 = 255.6546515665
	fmt.Println("Float32", smallFloat)
	fmt.Printf("Type of Varible is : %T \n", smallFloat)

	var smallFloat64 float64 = 255.6546515665
	fmt.Println("Float32", smallFloat64)
	fmt.Printf("Type of Varible is : %T \n", smallFloat64)

	// Default Values and some aliases

	var anotherVarible int
	fmt.Println("Another Varible", anotherVarible)
	fmt.Printf("Type of Varible is : %T \n", anotherVarible)

	var anotherVarible2 string
	fmt.Println("Another Varible", anotherVarible2)
	fmt.Printf("Type of Varible is : %T \n", anotherVarible2)

	// No datatype style
	// go automatically detects the datatype
	var website = "https://www.google.com"
	fmt.Println("Website", website)
	fmt.Printf("Type of Varible is : %T \n", website)

	// no var style

	numberOfUsers := 1003334500
	fmt.Println("Number of Users", numberOfUsers)

	fmt.Printf("Your login token is : %s \n", LoginToken)
	fmt.Printf("Type of login token  is : %T \n", LoginToken)

}
