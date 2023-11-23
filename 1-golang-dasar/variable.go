package main

import "fmt"

func main() {
	// membuat variable dengan deklarasi tipe data
	var name string

	name = "John Doe"
	fmt.Println(name)

	name = "Jane Doe"
	fmt.Println(name)

	// membuat variable tanpa deklarasi tipe data
	var age = 20

	fmt.Println(age)

	// membuat variable shorthand
	hair := "fluffy"
	fmt.Println(hair)

	hair = "curly"
	fmt.Println(hair)

	// membuat multiple variable
	var (
		firstName  = "John"
		middleName = "Wick"
		lastName   = "Doe"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
