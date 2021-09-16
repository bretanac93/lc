package sortedsquares

import "sort"

func mapSlice(arr []int, fn func(item int) int) []int {
	res := make([]int, 0)

	for _, item := range arr {
		res = append(res, fn(item))
	}

	return res
}

func sortedSquares(nums []int) []int {
	squares := mapSlice(nums, func(item int) int {
		return item * item
	})

	sort.Ints(squares)

	return squares
}
