package tribonacci

func tribonacci(n int) int {
	dp := make([]int, 40)

	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}
