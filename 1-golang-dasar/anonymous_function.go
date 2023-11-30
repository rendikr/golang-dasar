package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Access Denied,", name)
	} else {
		fmt.Println("Welcome,", name)
	}
}

func accessCheck(name string) bool {
	if name == "Jane" {
		return true
	}

	return false
}

func main() {
	blacklist := func(name string) bool {
		return name == "Jane"
	}

	registerUser("John", blacklist)
	registerUser("Jane", func(name string) bool {
		return name == "Jane"
	})
}
