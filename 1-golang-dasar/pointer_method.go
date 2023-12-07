package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	john := Man{"John"}
	john.Married()

	fmt.Println(john.Name) // Mr. John
}
