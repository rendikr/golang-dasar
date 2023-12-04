package main

import "fmt"

func endApps() {
	fmt.Println("End App")
	errorMessage := recover()
	fmt.Println("Error:", errorMessage)
}

func runApps(error bool) {
	defer endApps() // execute this function once all of this function logic has finished
	if error {
		panic("Error Happen!")
	}
}

func main() {
	runApps(true)
}
