package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp() // execute this function once all of this function logic has finished
	if error {
		panic("Error Happen!")
	}
}

func main() {
	runApp(true)
}
