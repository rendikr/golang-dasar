package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) { // pointer sebagai params input, maka value dari variable aslinya akan ikut berubah
	address.Country = "Indonesia"
}

func main() {
	var address = Address{}
	ChangeCountryToIndonesia(&address) // kirimkan reference sebagai params

	fmt.Println(address) // {  Indonesia}
}
