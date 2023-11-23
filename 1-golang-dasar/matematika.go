package main

import "fmt"

func main() {
	var a = 10
	var b = 2

	var tambah = a + b
	var kurang = a - b
	var kali = a * b
	var bagi = a / b
	var sisa = a % b

	fmt.Println(tambah)
	fmt.Println(kurang)
	fmt.Println(kali)
	fmt.Println(bagi)
	fmt.Println(sisa)

	// augmented assignments
	var i = 10

	i += 5 // i = i + 5
	fmt.Println(i)

	i -= 5 // i = i - 5
	fmt.Println(i)

	i *= 5 // i = i * 5
	fmt.Println(i)

	i /= 5 // i = i / 5
	fmt.Println(i)

	i %= 5 // i = i % 5
	fmt.Println(i)

	var n = 3

	n++ // n = n + 5
	fmt.Println(n)
	n-- // n = n - 5
	fmt.Println(n)
}
