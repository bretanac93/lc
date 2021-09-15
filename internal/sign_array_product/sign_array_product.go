package signarrayproduct

func filter(nums []int, condition func(num int) bool) []int {
	result := make([]int, 0)

	for _, item := range nums {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}

func arraySign(nums []int) int {
	zerosCount := filter(nums, func(num int) bool {
		return num == 0
	})

	if len(zerosCount) > 0 {
		return 0
	}

	negativesCount := filter(nums, func(num int) bool {
		return num < 0
	})

	if len(negativesCount)%2 == 0 {
		return -1
	}

	return 1
}
