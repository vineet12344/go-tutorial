package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // usually they are pointers pointers

var mut sync.Mutex // pointer

func main() {
	// fmt.Println("Welcome to go routines")
	// go greeter("HELLO")
	// greeter("WORLD")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Print("OOPS in Endpoint for %s", endpoint)
	} else {

		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)

	}
	fmt.Printf("Get request for %s completed \n", endpoint)
}

// func greeter(s string) {

// 	for range 6 {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)

// 	}
// }
