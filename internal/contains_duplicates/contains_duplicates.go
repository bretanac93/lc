package containsduplicates

func containsDuplicate(nums []int) bool {
	freq := make(map[int]int)

	for _, item := range nums {
		_, found := freq[item]

		if found {
			return true
		}

		freq[item] = 1
	}

	return false
}
