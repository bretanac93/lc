package romantointeger

func romanToInt(s string) int {

	dict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && dict[string(s[i])] < dict[string(s[i+1])] {
			sum -= dict[string(s[i])]
		} else {
			sum += dict[string(s[i])]
		}
	}

	return sum
}
