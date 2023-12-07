package main

import (
	"errors"
	"fmt"
)

func Divider(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Invalid num2")
	}

	return num1 / num2, nil
}

func main() {
	result, err := Divider(100, 10)
	if err == nil {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
