package main

import "fmt"

func main() {
	// pass by value
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // copy value

	address2.City = "Bandung"
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) // city berubah menjadi Bandung

	// pass by reference
	address3 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address4 := &address3 // reference

	address4.City = "Bandung"
	fmt.Println(address3) // city ikut berubah
	fmt.Println(address4) // city berubah menjadi Bandung
}
