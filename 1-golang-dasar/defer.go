package main

import "fmt"

func logging() {
	fmt.Println("Finished!")
}

func runApplication() {
	defer logging() // execute this function once all of this function logic has finished
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
