package main
import "fmt"

func main() {
	fmt.Println(" !!! Welcome to class on pointes !!! ")
	p := fmt.Println

	// var ptr *int
	// fmt.Println("Value of pointer is : ", ptr)

	myNumber := 23
	// p(&myNumber)
	// p(myNumber)

	var ptr = &myNumber
	// fmt.Println("Value of ptr is : ", ptr)
	// fmt.Println("Value of ptr is : ", *ptr)

	p("Value of ptr is : ", ptr)
	p("Value of ptr is : ", *ptr)

}
