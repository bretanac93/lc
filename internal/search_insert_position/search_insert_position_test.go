package searchinsertposition

import (
	"testing"
)

type testCase struct {
	nums     []int
	target   int
	expected int
}

func TestSearch(t *testing.T) {
	tests := []testCase{
		{
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			nums:     []int{1},
			target:   0,
			expected: 0,
		},
		{
			nums:     []int{-1, 1},
			target:   0,
			expected: 1,
		},
	}

	for _, tt := range tests {
		res := searchInsert(tt.nums, tt.target)

		if res != tt.expected {
			t.Errorf("searchInsert(%v, %d) got %d, expected %d", tt.nums, tt.target, res, tt.expected)
		}
	}
}
