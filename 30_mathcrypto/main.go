package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to Maths in Golang")

	// var numOne int = 2
	// var numTwo float64 = 4.5

	// ! fmt.Println("The Sum is : ", numOne+int(numTwo))

	// random number MATH and CRYPTO
	// rand.New(rand.NewSource(time.Now().UnixNano()))
	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)

}
