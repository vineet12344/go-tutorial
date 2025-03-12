package main

import "fmt"

func main() {
	// fmt.Println("Welocme to Functions in Golang ")

	// var result int = adder(4, 4)

	// fmt.Println(result)

	// fmt.Println("Profunction ")

	// result2, msg := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// fmt.Println(msg, result2)

	func() {
		fmt.Println("Hello froom anonymus function")
	}()

	var greetings func(string) = greet

	greetings("Vineey")

}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(nums ...int) (int, string) {
	var sum int

	for _, num := range nums {
		sum += num
	}

	return sum, "Added successfully "
}

func greet(name string) {
	fmt.Println("Hello Warald ", name)
}
