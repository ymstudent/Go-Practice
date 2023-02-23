package cn

//leetcode submit region begin(Prohibit modification and deletion)
// 由于机器人每次只能向下或者向右移动一步，所以当前位置的路径数 = 上方过来的路径数+左边过来的路径数
// 即 dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 对于第一行 dp[0][j]，或者第一列 dp[i][0]，由于都是在边界，所以只能为 1
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	// 整理边界条件
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1 // 令第一列的值都为1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1 // 令第一行的值都为1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 状态转移方程
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
