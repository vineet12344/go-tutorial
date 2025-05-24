package main

import (
	"fmt"
	"io"
	"net/http"
)

// import "http"

func main() {
	fmt.Println("HTTP SERVER")

	const url string = "https://httpbin.org/get"

	response, err := http.Get(url)

	checkError(err)

	fmt.Printf("Response is of Type: %T \n", response)
	defer response.Body.Close() // Callers Responsiblity !! :)

	databytes, err := io.ReadAll(response.Body)
	checkError(err)

	println("Data from API : ", string(databytes))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
