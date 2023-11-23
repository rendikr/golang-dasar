package main

import "fmt"

func main() {
	var names [3]string // buat array dengan 3 buah index dan tipe data string
	names[0] = "John"
	names[1] = "Jane"
	names[2] = "Doe"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{ // penggunaan [...] berarti jumlah index tidak ditentukan. namun penulisan values nya harus lgsg ditentukan
		90,
		80,
		85,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
	values[2] = 95
	fmt.Println(values[2])
}
