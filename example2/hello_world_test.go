package example2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := helloWorld("Mark")
		require.Equal(t, "Hello Mark", result, "result must be 'Hello Mark'")
	})

	t.Run("Test 2", func(t *testing.T) {
		result := helloWorld("Mark")
		require.Equal(t, "Hello Mark", result, "result must be 'Hello Mark'")
	})
}

// Test with multi condition using struct list
func TestTableTest(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Test1",
			request:  "Mark",
			expected: "Hello Mark",
		},
		{
			name:     "Test2",
			request:  "Zucker",
			expected: "Hello Zucker",
		},
		{
			name:     "Test3",
			request:  "Sundai",
			expected: "Hello Sundai",
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := helloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchMarkTable(b *testing.B) {
	benchmarks := []struct {
		name  string
		value string
	}{
		{
			name:  "Firstname",
			value: "Bos",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				helloWorld(benchmark.value)
			}
		})
	}
}
func BenchMarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloWorld("hallo")
	}
}
