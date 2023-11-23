package main

import "fmt"

func main() {
	fmt.Println("John")
	fmt.Println("Doe")
	fmt.Println("Jumlah Karakter:", len("John Doe"))
	fmt.Println("Karakter ke 0:", "John Doe"[0])
	fmt.Println("Karakter ke 5:", "John Doe"[5])
}
