package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke:", counter)
		counter++
	}

	fmt.Println("For Statement")

	// for statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke:", counter)
	}

	fmt.Println("For Range")

	// for range
	names := []string{"John", "Jane", "doe"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Nama:", names[i])
	}

	for index, name := range names {
		fmt.Println("Index:", index, "=", name)
	}

	fmt.Println("Selesai")
}
