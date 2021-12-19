package main

func RemoveDuplicates(nums []int) int {
	l := 1
	r := 1

	for r < len(nums) {
		if nums[r] != nums[r - 1] {
			nums[l] = nums[r - 1]
			l++
		}
		r++
	}
	return l
}
