package binarysearch

func binSearch(nums []int, target int, min, max int) int {
	if min > max {
		return -1
	}

	mid := (min + max) / 2

	if nums[mid] < target {
		return binSearch(nums, target, mid+1, max)
	} else if nums[mid] > target {
		return binSearch(nums, target, min, mid-1)
	} else {
		return mid
	}
}

func search(nums []int, target int) int {
	return binSearch(nums, target, 0, len(nums))
}
