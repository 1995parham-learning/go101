package fibonacci_test

import (
	"testing"

	"github.com/1995parham-learning/go101/fibonacci"
)

func TestFibonacci(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "2",
			input:    2,
			expected: 2,
		},
		{
			name:     "3",
			input:    3,
			expected: 3,
		},
		{
			name:     "4",
			input:    4,
			expected: 5,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if fibonacci.Fibonacci(tc.input) != tc.expected {
				t.Errorf("Fibonacci(%d) ==> %d != %d)", tc.input, fibonacci.Fibonacci(tc.input), tc.expected)
			}
		})
	}
}
