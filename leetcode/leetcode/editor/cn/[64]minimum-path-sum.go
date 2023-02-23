package cn

//leetcode submit region begin(Prohibit modification and deletion)
// 62题变化版
// 状态转移方程变为 dp[i][j] = grid[i][j] + minMinPathSum(dp[i-1][j], dp[i][j-1])
// 同样注意第一行和第一列的边界问题即可
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	ySum := 0
	for i := range dp {
		dp[i] = make([]int, n)
		ySum += grid[i][0]
		dp[i][0] = ySum
	}
	xSum := 0
	for j := 0; j < n; j++ {
		xSum += grid[0][j]
		dp[0][j] = xSum
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + minMinPathSum(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

func minMinPathSum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
