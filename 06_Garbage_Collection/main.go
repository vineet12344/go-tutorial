package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Printf("Current Process is using %d no of Go Routines \n", runtime.NumGoroutine())
	fmt.Printf("Current Process is using %d no of CPU's \n", runtime.NumCPU())
	fmt.Printf("Current Process is using %d compiler \n", runtime.NumCgoCall())

}
