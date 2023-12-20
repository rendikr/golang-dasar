package helper

import "testing"

func TestHelloWorldJohn(t *testing.T) {
	result := HelloWorld("John")

	if result != "Hello John" {
		t.Error("Result should be Hello John") // equal to t.Fail()
	}
}

func TestHelloWorldJane(t *testing.T) {
	result := HelloWorld("Jane")

	if result != "Hello Jane" {
		t.Fatal("Result should be Hello Jane") // equal to t.FailNow()
	}
}
