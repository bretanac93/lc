package validparentheses

import (
	"testing"

	"github.com/bretanac93/algo/internal/common"
)

func TestValidParentheses(t *testing.T) {
	testCases := []common.TestCase{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, tt := range testCases {
		res := isValid(tt.Input.(string))

		if res != tt.Expected.(bool) {
			t.Errorf("expected %v -- got %v", tt.Expected, res)
		}
	}
}
