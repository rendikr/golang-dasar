package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "John Doe"
	var encoded = base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println(string(decoded))
}
