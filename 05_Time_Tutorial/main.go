package main

import (
	"fmt"
	"time"

	"github.com/hako/durafmt"
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

	duration, err := durafmt.ParseStringShort("354h22m3.24s")
	// duration2, err := durafmt.Parse("354h22m3.24s").String()

	if err != nil {
		fmt.Println("Error in durafmt.ParseString()")
	} else {
		fmt.Println(duration)
		// fmt.Println(duration2)

	}

}
