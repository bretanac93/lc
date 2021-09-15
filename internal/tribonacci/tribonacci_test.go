package tribonacci

import "testing"

type testCase struct {
	input    int
	expected int
}

func TestTribonacci(t *testing.T) {
	testCases := []testCase{
		{
			input:    4,
			expected: 4,
		},
		{
			input:    25,
			expected: 1389537,
		},
	}

	for _, tt := range testCases {
		res := tribonacci(tt.input)
		if res != tt.expected {
			t.Errorf("expected %d -- got %d", tt.expected, res)
		}
	}
}
