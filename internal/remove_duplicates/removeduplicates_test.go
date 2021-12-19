package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := []int{0, 1, 2, 3, 4}

	count := RemoveDuplicates(arr)
	assert.Equal(t, 5, count)

	diff := false
	for i := 0; i < count; i++ {
		if arr[i] != expected[i] {
			diff = true
		}
	}

	assert.False(t, diff)
}
