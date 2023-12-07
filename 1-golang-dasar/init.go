package main

import (
	"1-golang-dasar/database"
	_ "1-golang-dasar/internal" // blank identifier (_) menandakan akan menggunakan package internal tapi tidak meng-eksekusi function init()
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
