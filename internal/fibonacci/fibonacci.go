package fibonacci

func fib(n int) int {
	dp := make([]int, 35)

	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
