package reverseinteger

import "math"

func reverse(x int) int {

	inverted := 0

	for x != 0 {
		digit := x % 10
		inverted = inverted*10 + digit

		x /= 10
	}

	if inverted > math.MaxInt32 || inverted < math.MinInt32 {
		return 0
	}

	return inverted
}
