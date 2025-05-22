package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://api.api-ninjas.com/v1/jokes"

func main() {
	fmt.Println("Welocme to URL Handling in GOlang")

	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
}
