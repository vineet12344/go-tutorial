package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "THursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days {    // It returns the index in i not the value
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	// fmt.Println("Day: ", day, "Index: ", index)
	// 	fmt.Printf("Index is and Day is %v \n", day)
	// }

	rogueValue := 0

	for rogueValue < 10 {

	// lab2:

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		if rogueValue == 2 {
			goto lco
		}

		fmt.Println("RogueValue is : ", rogueValue)
		rogueValue++
	}

lco:

	fmt.Println("Jumping at a label")
	// goto lab2
}
