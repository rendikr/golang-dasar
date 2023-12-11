package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("John Doe", "oh"))
	fmt.Println(strings.Split("John Doe", " "))
	fmt.Println(strings.ToLower("John Doe"))
	fmt.Println(strings.ToUpper("John Doe"))
	fmt.Println(strings.Trim("  John Doe  ", " "))
	fmt.Println(strings.ReplaceAll("John Doe", "John", "Jane"))
}
