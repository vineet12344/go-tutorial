package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Golang")
	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	// fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")

	case 2:
		fmt.Println("You can move 2 pos")
	case 3:
		fmt.Println("You can move 3 pos")
		fallthrough
	case 4:
		fmt.Println("You can move 4 pos")
	case 5:
		fmt.Println("You can move 5 pos")
		fallthrough
	case 6:
		fmt.Println("You can move 6 pos")
	default:
		fmt.Println("What was that ?")

	}

}
