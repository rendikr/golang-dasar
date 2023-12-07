package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// pass by reference
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // reference, address2 type is a pointer to Address

	address2.City = "Bandung"
	fmt.Println(address1) // {Bandung Jawa Barat Indonesia}
	fmt.Println(address2) // &{Bandung Jawa Barat Indonesia}

	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} // akan mengganti address2 value menjadi Address yang baru
	fmt.Println(address1)                                      // {Bandung Jawa Barat Indonesia}
	fmt.Println(address2)                                      // &{Jakarta DKI Jakarta Indonesia}

	// pass by reference
	address3 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address4 := &address3 // reference, address4 type is a pointer to Address

	address4.City = "Bandung"
	fmt.Println(address3) // {Bandung Jawa Barat Indonesia}
	fmt.Println(address4) // &{Bandung Jawa Barat Indonesia}

	*address4 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // akan mengganti reference value asli dari address4 (termasuk address3) menjadi Address yang baru
	fmt.Println(address3)                                      // {Jakarta Jawa Barat Indonesia}
	fmt.Println(address4)                                      // &{Jakarta DKI Jakarta Indonesia}
}
