package romantointeger

import (
	"testing"

	"github.com/bretanac93/algo/internal/common"
)

func TestRomanToInteger(t *testing.T) {
	testCases := []common.TestCase{
		{"III", 3},
		{"IV", 4},
		{"VI", 6},
		{"MCMXCIII", 1993},
		{"MCMXCIV", 1994},
	}

	for _, tt := range testCases {
		res := romanToInt(tt.Input.(string))
		if res != tt.Expected {
			t.Errorf("expected: %v -- got %v", tt.Expected, res)
		}
	}
}
