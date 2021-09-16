package sortedsquares

import (
	"testing"

	"github.com/bretanac93/algo/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	testCases := []common.TestCase{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, tt := range testCases {
		res := sortedSquares(tt.Input.([]int))

		assert.ElementsMatch(t, res, tt.Expected)
	}
}
