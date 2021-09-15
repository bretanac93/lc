package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchFoundElement(t *testing.T) {
	input := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expectedOutput := 4
	result := search(input, target)

	assert.Equal(t, expectedOutput, result)
}

func TestSearchNotFoundElement(t *testing.T) {
	input := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	expectedOutput := -1
	result := search(input, target)

	assert.Equal(t, expectedOutput, result)
}
