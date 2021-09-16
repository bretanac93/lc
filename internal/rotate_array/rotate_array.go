package rotatearray

func reverse(arr []int, start, end int) {
	for start < end {
		tmp := arr[start]
		arr[start] = arr[end]
		arr[end] = tmp

		start++
		end--
	}
}

func rotate(nums []int, k int) {
	c := len(nums)
	k %= c

	if k == 0 {
		return
	}

	reverse(nums, 0, c-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, c-1)
}
