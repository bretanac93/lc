package findinganagrams

import (
	"strconv"
)

func stringHash(str string) string {
	freq := make([]int, 26)

	for _, letter := range str {
		idx := letter - 'a'
		freq[idx]++
	}

	res := ""

	for _, idx := range freq {
		res += "#" + strconv.Itoa(idx)
	}

	return res
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	res := make([][]string, 0)

	for _, str := range strs {
		k := stringHash(str)
		if _, ok := groups[k]; !ok {
			groups[k] = []string{str}
		} else {
			groups[k] = append(groups[k], str)
		}
	}

	for _, vals := range groups {
		res = append(res, vals)
	}
	return res
}
