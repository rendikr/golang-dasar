package main

import "fmt"

func getFullName() (string, string) {
	return "John", "Doe"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// ignore the "lastName" returned value
	myName, _ := getFullName()
	fmt.Println(myName)
}
