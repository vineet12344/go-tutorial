package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in GO Lang Tutorial")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Printf("TYpe of Map is : %T \n", languages)
	fmt.Println("JS is for ", languages["JS"])
	fmt.Println("PY is for ", languages["Py"])

	// delete(languages, "Py")

	fmt.Println("All languages are: \n \n ####################################### \n", languages)

	// Loops are mfs in golang

	for key, value := range languages {

		fmt.Print("[", key, "] -> ")
		fmt.Print("[", value, "] \n")

	}

	for _, value := range languages {
		// fmt.Println(key, value)
		fmt.Printf(" value is %v \n", value)

	}

}
