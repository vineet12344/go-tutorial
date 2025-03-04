package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()

	// fmt.Println("Current time is ", currentTime)
	fmt.Println("Day: ", currentTime.Day())
	fmt.Println("Month: ", currentTime.Month())
	fmt.Println("Year : ", currentTime.Year())

	fmt.Println(currentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2003, time.November, 12, 23, 23, 0, 0, time.UTC)
	// fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))

}
