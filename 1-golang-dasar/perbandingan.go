package main

import "fmt"

func main() {
	// operasi perbandingan
	var name1 = "John"
	var name2 = "John"

	var result bool = name1 == name2
	fmt.Println(result)

	result = name1 != name2
	fmt.Println(result)

	var a = 4
	var b = 6
	result = a > b
	fmt.Println(result)

	result = a < b
	fmt.Println(result)
}
