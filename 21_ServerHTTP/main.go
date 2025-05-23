package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("WELOCME TO FRONTEND REQUEST MAKER (i guess)")
	// PerformGetRequest()
	PerformPostJsonRequest()

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformGetRequest() {
	const url string = "http://localhost:8000/get"

	response, err := http.Get(url)

	checkError(err)

	defer response.Body.Close()

	fmt.Println("STATUS CODE: ", response.StatusCode)
	fmt.Println("CONTENT LENGHT: ", response.ContentLength)

	var responseString strings.Builder

	content, err := io.ReadAll(response.Body)
	checkError(err)

	length, err := responseString.Write(content)
	checkError(err)
	fmt.Println("responseString length: ", length)
	// fmt.Println("CONTENT LENGHT: ", string(content))

	fmt.Println("Response String: ", responseString.String())

}

func PerformPostJsonRequest() {
	const myUrl string = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
		"CourseName": "Lets GO with golang :)"
	}
	`)

	resp, err := http.Post(myUrl, "Application/json", requestBody)
	checkError(err)

	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	checkError(err)

	fmt.Println("POST REQUEST RESPONSE IS: ", string(content))

}
