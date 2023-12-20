package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("John")
	assert.Equal(t, "Hello John", result, "Result should be Hello John") // will call t.Fail() on test fail
	fmt.Println(">>> [TestHelloWorldAssertion] Finished...")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("John")
	require.Equal(t, "Hello John", result, "Result should be Hello John") // will call t.FailNow() on test fail
	fmt.Println(">>> [TestHelloWorldRequire] Finished...")
}
