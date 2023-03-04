package cn

//leetcode submit region begin(Prohibit modification and deletion)

// dp[i] = max(dp[i], dp[i-j*j] + 1)
func numSquares(n int) int {
	dp := make([]int, n+1) // 注意：n+1
	for i := 1; i <= n; i++ {
		dp[i] = i // 当前数字的最大结果，eg：i=4, 最坏结果为 4=1+1+1+1 即为 4 个数字
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
