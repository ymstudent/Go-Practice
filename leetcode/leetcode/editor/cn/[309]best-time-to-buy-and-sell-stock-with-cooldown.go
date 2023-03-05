package cn
//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		if i-2== -1 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
			dp[i][1] = max(dp[i-1][1], - prices[i])
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
	}

	return dp[n-1][0]
}
//leetcode submit region end(Prohibit modification and deletion)
