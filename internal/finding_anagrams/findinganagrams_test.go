package findinganagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	testCase := []string{"word", "sword", "drow", "rowd", "iced", "dice"}
	expectedRes := [][]string{
		{"word", "drow", "rowd"},
		{"sword"},
		{"iced", "dice"},
	}
	res := groupAnagrams(testCase)

	for _, resItem := range expectedRes {
		assert.Contains(t, res, resItem)
	}
}
