package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welocme to File in Golang !!")
	content := "Vineet Salve"
	file, err := os.Create("./MyName.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile(file.Name())

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkError(err)

	fmt.Println("Data Read From File: ", string(databyte))

}

func checkError(err error) {
	if err != nil {
		panic("Error Occured")
	}
}
