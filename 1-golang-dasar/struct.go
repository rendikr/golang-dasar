package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var john Customer
	john.Name = "John"
	john.Address = "JL. ABC"
	john.Age = 20

	fmt.Println(john)
	fmt.Println(john.Name)
	fmt.Println(john.Address)
	fmt.Println(john.Age)

	jane := Customer{
		Name:    "Jane",
		Address: "JL. ABC",
		Age:     21,
	}
	fmt.Println(jane)

	peter := Customer{"Peter", "JL. ABC", 22}
	fmt.Println(peter)
}
