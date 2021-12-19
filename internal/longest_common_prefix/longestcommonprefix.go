package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		ver := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || ver != strs[j][i] {
				return prefix
			}
		}
		prefix += string(ver)
	}

	return prefix
}
