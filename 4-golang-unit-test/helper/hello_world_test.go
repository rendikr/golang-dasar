package helper

import (
	"fmt"
	"runtime"
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

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unable to run unit test on Mac OS")
	}

	result := HelloWorld("John")
	require.Equal(t, "Hello John", result, "Result should be Hello John") // will call t.FailNow() on test fail
}

func TestMain(m *testing.M) {
	// before running unit test
	fmt.Println("Before Unit Test")

	m.Run() // execute all unit test

	// after running unit test
	fmt.Println("After Unit Test")
}

func TestSubTest(t *testing.T) {
	fmt.Println(">>> [TestSubTest] Started...")
	t.Run("John", func(t *testing.T) {
		result := HelloWorld("John")
		require.Equal(t, "Hello John", result, "Result should be Hello John")
	})
	t.Run("Jane", func(t *testing.T) {
		result := HelloWorld("Jane")
		require.Equal(t, "Hello Jane", result, "Result should be Hello Jane")
	})
	fmt.Println(">>> [TestSubTest] Finished...")
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(John)",
			request:  "John",
			expected: "Hello John",
		},
		{
			name:     "HelloWorld(Jane)",
			request:  "Jane",
			expected: "Hello Jane",
		},
		{
			name:     "HelloWorld(Peter)",
			request:  "Peter",
			expected: "Hello Peter",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorldJohn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("John")
	}
}

func BenchmarkHelloWorldJane(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Jane")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("John", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("John")
		}
	})
	b.Run("Jane", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jane")
		}
	})
}
