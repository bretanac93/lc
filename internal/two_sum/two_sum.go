package twosum

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)

	found := make([]int, 2)

	for idx, item := range nums {
		diff := target - item

		foundIdx, exists := dict[diff]

		if !exists {
			dict[item] = idx
		} else {
			found[0] = foundIdx
			found[1] = idx
		}
	}

	return found
}
