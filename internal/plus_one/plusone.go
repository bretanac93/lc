package plus_one

func invertArray(arr []int) []int {
	res := make([]int, len(arr))

	for i := len(arr); i > 0; i-- {
		res = append(res, arr[i])
	}

	return res
}

func plusOne(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 9 {
			arr[i] = 0
		} else {
			arr[i]++
			return arr
		}
	}
	arr = append([]int{1}, arr...)
	return arr
}
