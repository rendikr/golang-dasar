package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["age"]) // karena key belum di-assign value nya, maka value nya akan menggunakan default value utk tipe data tersebut

	book := make(map[string]string) // alternatif lain membuat map baru
	book["title"] = "Buku Golang"
	book["author"] = "John Doe"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups") // menghapus key "ups" pada map

	fmt.Println(book)

}
