package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welocme to JSON Form Post Request")

	sendPostFormRequest()
}

func sendPostFormRequest() {
	const myURL string = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("Name", "Vineet Salve")
	data.Add("Class", "D")
	data.Add("CGPA", "7.5")

	resp, err := http.PostForm(myURL, data)
	checkError(err)

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	checkError(err)

	var responseBuilder strings.Builder

	length, err := responseBuilder.Write(content)
	checkError(err)

	fmt.Println("Form data Length is: ", length)
	fmt.Println("Form data received is: ", responseBuilder.String())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
