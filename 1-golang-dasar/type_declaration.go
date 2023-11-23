package main

import "fmt"

func main() {
	type NoKTP string

	var ktpJohn NoKTP = "111111111"
	fmt.Println(ktpJohn)
	fmt.Println(NoKTP("222222222"))
}
