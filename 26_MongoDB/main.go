package main

import (
	"fmt"
	"log"
	"mongodbTutorial/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is Started .....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at Port 4000....")

}
