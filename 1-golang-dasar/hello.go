package main

import (
	"1-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("John")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)            // returns undefined
	// fmt.Println(helper.sayGoodbye("John")) // returns undefined
}
