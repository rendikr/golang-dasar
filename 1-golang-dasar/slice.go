package main

import "fmt"

func main() {
	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice1 := months[4:7]
	fmt.Println(slice1) // ["Mei", "Juni", "Juli"]

	slice2 := months[:4]
	fmt.Println(slice2) // ["Januari", "Februari", "Maret", "April"]

	slice3 := months[7:]
	fmt.Println(slice3) // ["Agustus", "September", "Oktober", "November", "Desember"]

	slice4 := months[:]
	fmt.Println(slice4) // ["Januari", ..., "Desember"]

	// fungsi dalam slice
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1)) // capacity dari pointer awal sampai akhir dari array

	monthSlice := append(slice1, "Agustus")
	fmt.Println(monthSlice)
	monthSlice[0] = "Ups"
	fmt.Println(monthSlice)
	fmt.Println(months)

	daySlice := make([]string, 2, 5) // membuat slice baru dengan jumlah index 2 & capacity 5
	daySlice[0] = "Senin"
	daySlice[1] = "Selasa"
	fmt.Println(daySlice)

	// copy slice
	fromSlice := months[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
