package main

import "fmt"

// "fmt"

func main() {
	// var fruitList = []string{"Apple", "Banana", "Coconut"}

	// fmt.Printf("Type of Slice is : %T \n", fruitList)
	// fmt.Println("Fruit List", fruitList)

	// fruitList = append(fruitList, "Peach", "Strawberry", "Mango ")

	// fmt.Println("After updating the Slice we get :", fruitList)

	// fmt.Println("Length of this Slice is : ", len(fruitList))

	// fmt.Println(append(fruitList[1:]))

	// highScores := make([]int, 4)
	// // fmt.Println(highScores)
	// highScores[0] = 234
	// highScores[1] = 945
	// highScores[2] = 456
	// highScores[3] = 589
	// // highScores[4] = 398

	// // fmt.Println(highScores)
	// highScores = append(highScores, 234, 356, 758)
	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// // fmt.Println("Sorted Array : ", highScores)
	// // fmt.Println(sort.IntsAreSorted(highScores))

	// var nums = []int{4, 3, 1, 6, 9, 7, 5}
	// sort.Ints(nums)
	// // println(nums)
	// fmt.Println(nums)

	// var names = []string{"Charlie", "Bob", "Alan"}
	// sort.Strings(names)
	// fmt.Println(names)
	// fmt.Println(sort.SearchStrings(names, "Vineet"))

	// var ages = []int{20, 17, 19, 23, 21, 25}
	// // sort.Reverse(sort.Ints(ages))

	// fmt.Println("Sorted in reverse order: ", ages)

	// removing an element in slice based on the index
	// Also simulating the database opreations

	// slice of courses
	// var courses = []string{"reactjs", "javascript", "swift", "Python", "Ruby"}

	// fmt.Println(courses)

	// var index int = 2

	// // Super important
	// courses = append(courses[:index], courses[index+1:]...)
	// fmt.Println(courses)

	// ------------------------------------------------------------------------------------------------------------------

	// Q1. Create an array of 5 integers and print it
	var nums = []int{1, 2, 3, 4, 5}
	fmt.Println("Nums are :", nums)

	// Q2. Sum all the elements in the array and return the result

	fmt.Println("Printing sum using a Function: ", arraySum(nums))

	// Copy one array to another and print both
	var array1 = []int{3, 5, 7, 80, 2, 4}
	var array2 = []int{}

	fmt.Println("Length of array1 is : ", len(array1))

	copyArray(array1, array2)

}

func arraySum(nums []int) int {
	var sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

func copyArray(array1 []int, array2 []int) {

	array2 = append(array2, array1...)

	fmt.Println("Array1: ", array1)
	fmt.Println("Array2 :", array2)

}
