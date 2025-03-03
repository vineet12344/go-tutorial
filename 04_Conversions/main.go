package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please rate our App")
	fmt.Println("Please type your rating: ")

	inputReader := bufio.NewReader(os.Stdin)

	rating, _ := inputReader.ReadString('\n')

	fmt.Println("THanks for the rating of : ", rating)

	// numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	numRating, err := strconv.ParseUint(strings.TrimSpace(rating), 2, 64)

	fmt.Println(numRating)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating : ", numRating+1)
	}

	floatNumRating, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	fmt.Println(floatNumRating + 0.5)

}
