package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data["name"])
	}
	// above code will output "Data Kosong"

	data2 := NewMap("hello")
	if data2 == nil {
		fmt.Println("Data2 Kosong")
	} else {
		fmt.Println(data2["name"])
	}
	// above code will output data2 contents
}
