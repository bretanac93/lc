package reverseinteger

import (
	"testing"

	"github.com/bretanac93/algo/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {

	testCases := []common.TestCase{
		// {12, 21},
		// {10, 1},
		// {223, 322},
		// {333, 333},
		{-123, -321},
	}

	for _, tt := range testCases {
		res := reverse(tt.Input.(int))

		assert.Equal(t, tt.Expected, res)
	}
}
