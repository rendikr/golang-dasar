package main

import "fmt"

func main() {
	const firstName string = "John"
	const lastName = "Doe"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// mengganti value dari variable di bawah akan menghasilkan error, dikarenakan constant merupakan variable yg value nya tidak bisa diubah lagi
	firstName = "Jane"
}
