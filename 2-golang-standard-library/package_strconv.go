package main

import (
	"fmt"
	"strconv"
)

func main() {
	// convert string to bool
	resultBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultBool)
	}

	// convert string to int
	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	// convert int to string
	resultString := strconv.Itoa(999)
	fmt.Println(resultString)
}
