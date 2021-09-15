package firstbadversion

func isBadVersion(version int) bool {
	return version == 5 || version == 4
}

func firstBadVersion(n int) int {
	min, max := 1, n

	for min <= max {
		mid := (max + min) / 2

		if !isBadVersion(mid) {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return min
}
