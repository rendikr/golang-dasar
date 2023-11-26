package main

import "fmt"

func main() {
	name := "John"

	switch name {
	case "John":
		fmt.Println("Hello, John")
	case "Jane":
		fmt.Println("How are you, Jane?")
	default:
		fmt.Println("Hello, stranger")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah oke")
	}
}
