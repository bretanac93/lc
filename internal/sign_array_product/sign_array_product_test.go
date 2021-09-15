package signarrayproduct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraySign(t *testing.T) {
	input := []int{-1, -2, -3, -4, 3, 2, 1}

	result := arraySign(input)

	assert.Equal(t, -1, result)
}
