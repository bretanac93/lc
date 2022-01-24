package str_str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	haystack := "c"
	needle := "c"
	expected := 0

	pos := strStr(haystack, needle)

	assert.Equal(t, expected, pos)
}

func naiveIndexOf(text, pattern string) int {
	n := len(text)
	m := len(pattern)

	for s := 0; s < (n - m) + 1; s++ {
		window := text[s : s+m]
		if pattern == window {
			return s
		}
	}
	return -1
}

func strStr(haystack, needle string) int {
	return naiveIndexOf(haystack, needle)
}
