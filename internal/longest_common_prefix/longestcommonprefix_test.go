package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	input := []string{"flower", "flow", "flight"}
	expectedOutput := "fl"

	output := longestCommonPrefix(input)
	assert.Equal(t, expectedOutput, output)
}

