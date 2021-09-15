package searchinsertposition

func search(nums []int, target int, min, max int) int {

	mid := (min + max) / 2

	if nums[mid] < target && mid < max {
		return search(nums, target, mid+1, max)
	} else if nums[mid] > target && mid > min {
		return search(nums, target, min, mid-1)
	} else if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return mid + 1
	} else {
		return mid
	}
}

func searchInsert(nums []int, target int) int {
	return search(nums, target, 0, len(nums)-1)
}
