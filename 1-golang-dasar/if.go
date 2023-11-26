package main

import "fmt"

func main() {
	name := "Jane"

	if name == "Doe" {
		fmt.Println("Hello, John!")
	} else if name == "Jane" {
		fmt.Println("How are you, Jane?")
	} else {
		fmt.Println("We're stranger")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah oke")
	}
}
