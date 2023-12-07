package main

import "fmt"

func main() {
	// var alamat1 = &Address{} // membuat alamat1 menjadi pointer ke value Address dengan data kosong
	var alamat1 = new(Address) // membuat alamat1 menjadi pointer ke value Address dengan data kosong. Alternatif dari &Address{}
	var alamat2 = alamat1      // alamamt2 otomatis pointer ke alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1) // &{  Indonesia}
	fmt.Println(alamat2) // &{  Indonesia}
}
