package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("John")

	if result != "Hello John" {
		panic("Result is not 'Hello John'")
	}
}
