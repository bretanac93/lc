package palindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := x
	inverted := 0

	for num != 0 {
		digit := num % 10
		inverted = inverted*10 + digit

		num /= 10
	}

	return inverted == x
}
