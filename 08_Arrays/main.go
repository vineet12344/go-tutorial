package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Go")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is:", fruitList)
	fmt.Println("length of Array", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Veggie List is : ", len(vegList))
}
