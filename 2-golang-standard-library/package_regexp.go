package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("john")) // returns true because string match regex (all lowercase)
	fmt.Println(regex.MatchString("jane"))
	fmt.Println(regex.MatchString("Doe")) // returns false because string doesn't match regex (has uppercase)

	fmt.Println(regex.FindAllString("john jane Doe", 10))
}
