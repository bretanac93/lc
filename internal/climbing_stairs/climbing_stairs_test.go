package climbingstairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	res := climbStairs(6)

	assert.Equal(t, 13, res)
}
