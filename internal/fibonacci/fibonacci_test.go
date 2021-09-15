package fibonacci

import "testing"

type testCase struct {
	input    int
	expected int
}

func TestFibonacci(t *testing.T) {
	testCases := []testCase{
		{
			input:    2,
			expected: 1,
		},
		{
			input:    3,
			expected: 2,
		},
		{
			input:    4,
			expected: 3,
		},
	}

	for _, tt := range testCases {
		res := fib(tt.input)

		if res != tt.expected {
			t.Errorf("got %d -- expected %d", tt.input, tt.expected)
		}
	}
}
