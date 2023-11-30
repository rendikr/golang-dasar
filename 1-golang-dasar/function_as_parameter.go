package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func sayHelloWithFilterAndFunctionType(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "shit" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("John", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("shit", filter)

	sayHelloWithFilterAndFunctionType("Doe", spamFilter)
}
