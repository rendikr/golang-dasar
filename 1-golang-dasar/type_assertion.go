package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result = random()
	// var resultString = result.(string)
	// fmt.Println(resultString)

	// var resultInt = result.(int)
	// fmt.Println(resultInt) // will return panic

	switch value := result.(type) {
	case string:
		fmt.Println("String:", value)
	case int:
		fmt.Println("Number:", value)
	default:
		fmt.Println("Unknown:", value)
	}
}
