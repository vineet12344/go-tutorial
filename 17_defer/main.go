package main

import "fmt"

func main() {
	defer fmt.Println("WORLD")
	defer fmt.Println("ONE")
	defer fmt.Println("TWO")
	fmt.Println("HELLO ")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
