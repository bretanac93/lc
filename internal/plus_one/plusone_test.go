package plus_one

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bretanac93/algo/internal/common"
)

func TestPlusOne(t *testing.T) {
	testCases := []common.TestCase{
		{
			Input:    []int{1, 0},
			Expected: []int{1, 1},
		},
		{
			Input:    []int{9},
			Expected: []int{1, 0},
		},
		{
			Input:    []int{1, 0, 3},
			Expected: []int{1, 0, 4},
		},
		{
			Input:    []int{9, 9, 9, 9},
			Expected: []int{1, 0, 0, 0, 0},
		},
	}

	for _, tt := range testCases {
		assert.Equal(t, tt.Expected, plusOne(tt.Input.([]int)))
	}
}
